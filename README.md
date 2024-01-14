# NET-CAT

## OVERVIEW

This project is a recreation of the NetCat utility in a server-client architecture. It can run in a server mode on a specified port, listening for incoming conenctions, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

NetCat, also known as nc, is a command line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, such as opening TCP connections, sending UDP packets, listening on arbitrary TCP and UDP ports, and port scanning.

## FEATURES

* TCP connection between server and multiple clients (relation of 1 to 10).
* Name requirement for clients.
* Control of connection quantity.
* Ability for clients to send messages to the chat.
* Filtering of empty messages from clients.
* Identification of messages sent by the time they were sent and the username of who sent them.
* Uploading of previous messages sent to the chat to new clients when they join.
* Notification of all clients when a new client joins or leaves the chat.
* Receipt of messages sent by other clients by all clients.
* Continued operation of other clients when a client leaves the chat.
* Default port setting if not port is specified
* Profanity filter that scans each message sent by clients for slurs/profanity.
* Any slurs detected sent by a client, he/she is kicked from the server.
* Any client who is kicked from the server, their name will be banned from joining again.

## USAGE

To start the TCP server, simply run the below command:

```
go build
```

```
go run main.go $port

```

```
./TCPChat $port

```

To run on default port (8989):

```
go run main.go

```

```
./TCPChat

```

To connect from client side, open new terminal and run below command:

```
nc localhost $port

```
Make sure the port number is exactly the same one server is on!

## AUTHORS

* Hussain
* Syed Hisham
* Abdulrahman

## CONCLUSION

This NetCat project is a simple but powerful tool for creating and managing network chats. It is easy to use and can be custommized to meet specific needs.

## LICENSES

This program developed within the scope of Reboot.