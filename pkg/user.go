package pkg

import (
	"net"
)

// get the name for the client
func GetUserName(conn net.Conn) string {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		PrintError("error reading user connection")
		return ""
	}
	name := string(buf[:n-1])
	// if name == "" {
	// 	return "Anonymous"
	// }
	// avoid last index because it is '\n'
	return name
}

// check if the user exist or not
func IsUserExist(u string, clients *map[string]Client) bool {
	for _, otherClient := range *clients {
		if u == otherClient.UserName {
			return true
		}
	}
	return false
}

// get ip
func GetClientIP(conn net.Conn) string {
	host, _, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		return ""
	}
	return host
}
