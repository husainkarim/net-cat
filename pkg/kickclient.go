package pkg

import (
	"fmt"
	"time"
)

// kick the client from the server
func KickClient(usrName string, clients *map[string]Client, messages chan string) {
	if client, ok := (*clients)[usrName]; ok {
		err := AddUsrToBanList(usrName)
		if err != nil {
			PrintError("error adding user to ban list")
		}
		delete(*clients, usrName)
		client.Conn.Write([]byte("\033[31mYou have been kicked from the server.\033[0m\n"))
		client.Conn.Close()

		//Notify other clients
		kickMsg := fmt.Sprintf("\033[36m[%s] (%s) has been kicked from the server.\033[0m\n", time.Now().Format("2006-01-02 15:04:05"), usrName)
		messages <- kickMsg
	}
}
