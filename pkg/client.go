package pkg

import "net"

type Client struct {
	UserName string
	Message  string
	DateTime string
	IPAddr   string
	Conn     net.Conn
}
