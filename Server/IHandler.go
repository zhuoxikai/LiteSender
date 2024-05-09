package main

import "net"

type IHandler interface {
	Handle(conn net.Conn)
}
