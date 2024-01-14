// ADDED BY SAYED:
// LINES: 26, 71, 112, 117-123, 132

package main

import (
	"TCPChat/pkg"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// created a map list to keep track of connected clients
//var activeClients map[string]Client = make(map[string]Client)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8989"
	CONN_TYPE = "tcp"
)

const DefaultChannel = "general"

var LogChat []string

//var bannedClients = pkg.LoadBannedClients("banned-clients.txt")

func main() {
	pkg.Logfile("", "log.txt")
	args := os.Args[1:]
	// Create a channel to send messages between goroutines.
	messages := make(chan string)
	// Create a TCP server.
	port := CONN_PORT // default port
	if len(args) > 0 {
		if len(args) != 1 { // check if they identify the port
			pkg.PrintUsage()
			return
		} else {
			port = args[0]
		}
	}
	// run the tcp chat at the identify port
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	listener = pkg.LimitListener(listener, 10)
	// defer listener.Close()
	fmt.Println("Server Start at PORT: " + port)
	// Create a wait group to limit the number of clients
	//clientsWaitGroup = sync.WaitGroup{}
	// Create a map to store the client information.
	clients := make(map[string]pkg.Client, 100)
	// Accept connections from clients.

	for {
		// accept connection from client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection.
		go HandleConnection(conn, &clients, messages)
	}
}

func HandleConnection(conn net.Conn, clients *map[string]pkg.Client, messages chan string) {
	// Send the welcome message.
	conn.Write(pkg.WelcomeMessage())
	// Get the client's username.
	userName := pkg.GetUserName(conn)
	bannedList, err := pkg.ReadBannedClientName("banned-clients.txt")
	if err != nil {
		pkg.PrintError("error reading banned clients list")
	}
	// check if the client exist or not
	if _, exists := (*clients)[userName]; exists {
		conn.Write([]byte("\033[1;33mThe name you entered already exist. Kindly try to connect with another name!\033[0m\n"))
		conn.Close()
		return
	}
	// check if the client in banned list or not
	for _, name := range bannedList {
		if name == userName {
			conn.Write([]byte("\033[1;33mYou have been banned from joining the server!\033[0m\n"))
			conn.Close()
			return
		}
	}

	channel := ""

	if channel == "" {
		channel = DefaultChannel
	}

	// create var for the client
	client := pkg.Client{UserName: userName}
	client.Conn = conn
	clientIP := pkg.GetClientIP(client.Conn)
	client.IPAddr = clientIP
	client.DateTime = time.Now().Format("2006-01-02 15:04:05")
	(*clients)[userName] = client
	// set join message for each client
	joinChat := "\033[1;32m[" + client.UserName + "] Has Joined Our Chat...\033[0m\n"
	// add the message to the Log array
	for _, otherClient := range *clients { // print the join message in others client chats
		if client.UserName != otherClient.UserName {
			otherClient.Conn.Write([]byte(joinChat))
		}
	}
	// print history to client
	for _, msg := range LogChat {
		client.Conn.Write([]byte(msg))
	}
	LogChat = append(LogChat, joinChat)
	// Start a goroutine to read messages from the client.
	go func() {
		// goroutine to hundle any client left the chat
		defer func() {
			delete(*clients, userName)

			var leftMsg string = fmt.Sprintf("\033[1;31m[%s] Has Left The Chat.\033[0m\n", userName)
			for _, otherClient := range *clients { // print the message to others in the chat
				otherClient.Conn.Write([]byte(leftMsg))
			}
			// add the left message to log
			LogChat = append(LogChat, leftMsg)
			fileText := strings.Join(LogChat, "\n")
			// update the log file
			pkg.Logfile(fileText, "log.txt")
			conn.Close()
		}()
		// set scanner in conn server to read all message sent by client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			if CheckMessage(message) { // if the message sent no empty proceed if not it will not be sent to others
				t := time.Now()
				msgWords := strings.Split(message, " ")
				// check for flag
				if strings.HasPrefix(message, "--") {
					Request := strings.Split(message, " ")
					if len(Request) != 2 { // handle len error
						er := "\"" + message + "\" bad request\n"
						client.Conn.Write([]byte(er))
						continue
					} else if strings.Contains(message, "--name") { // to change the name
						var exists bool = false
						for _, otherclient := range *clients {
							if Request[1] == otherclient.UserName {
								exists = true
							}
						}
						if exists {
							message = fmt.Sprintf("\033[1;33mName [%s] already exist. Kindly try different name!\033[0m\n", Request[1])
						} else { // if everything ok change the name
							oldname := userName
							userName = Request[1]
							client.UserName = userName
							(*clients)[oldname] = client
							message = fmt.Sprintf("[%s] has change his name to [%s]\n", oldname, userName)
						}
					} else if strings.Contains(message, "--ip") {
						//get ip address of client name passed after flag
						choseName := Request[1]
						clientIP = pkg.GetClientIP(client.Conn)
						if choseName == userName {
							client.Conn.Write([]byte("Entered name " + choseName + " is your registered name.\n"))
							continue
						} else if pkg.IsUserExist(choseName, clients) { // get the ip for the requested user
							message = fmt.Sprintf("Client: [%s]\nIP Address: [%s]\n", choseName, clientIP)
							client.Conn.Write([]byte(message))
							continue
						} else { // handle for not found client
							client.Conn.Write([]byte(choseName + " is not a valid active client name.\n"))
							continue
						}
					} else { // handle for wrong request
						er := "\"" + message + "\" bad request\n"
						client.Conn.Write([]byte(er))
						continue
					}
				} else { // if no flag
					message = fmt.Sprintf("[%s][%s]: %s\n", t.Format("2006-01-02 15:04:05"), userName, message)
				}
				//check if message contains any profanity
				if !pkg.NoProfanityInMsg(msgWords) {
					pkg.KickClient(userName, clients, messages)
				} else {
					if len(message) > 140 {
						conn.Write([]byte("\033[33mMaximum characters reached.\033[0m\n"))
					} else {
						// update the log
						LogChat = append(LogChat, message)
						messages <- message
					}
				}
			}
		}
	}()

	// Start a goroutine to write messages to the client.
	go func(client pkg.Client) {
		if client.Conn == nil { // handle if the connection lost
			fmt.Println(client.UserName + " lost connection")
			return
		}
		for {
			// Read a message from the channel.
			message := <-messages
			fileText := strings.Join(LogChat, "\n")
			// update log file
			pkg.Logfile(fileText, "log.txt")
			// print message to other
			for _, otherClient := range *clients {
				if !strings.Contains(message, otherClient.UserName) { // print the message to other
					otherClient.Conn.Write([]byte(message))
				}
			}
		}
	}(client)
}

func CheckMessage(list string) bool {
	for _, v := range list {
		if v > 32 && v < 127 {
			return true
		}
	}
	return false
}
