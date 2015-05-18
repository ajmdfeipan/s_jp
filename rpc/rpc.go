package rpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type State int

func (l *State) GetLine(line []byte, ack *bool) error {
	fmt.Println(string(line))
	return nil
}

func Start() {
	addr := "0.0.0.0:1970" //fmt.Sprintf("%s:%d", g.Config().Rpc.Addr, g.Config().Rpc.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr fail: %s", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail: %s", addr, err)
	}

	rpc.Register(new(State))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept occur error: %s", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
