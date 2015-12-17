package jsonrpc

import (

"fmt"
"net"
"net/rpc"
"net/rpc/jsonrpc"

)


type CC int

func (t *CC) Echo(args *string, reply *string) error {
    fmt.Println("Arg passed: " + *args)
    *reply = ">" + *args + "<"
    return nil
}

func Start() {

    arith := new(CC)
    rpc.Register(arith)

    tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12981")
    if err != nil {
        fmt.Println(err)
    }
    listener, err := net.ListenTCP("tcp", tcpAddr)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go jsonrpc.ServeConn(conn)
    }
}