package main

import (
	"./cron"
	"./http"
	"./rpc"
    "./jsonrpc"
	"log"
)

func main() {
	log.Println("staring...")
	go cala.Timeshow()
    go jsonrpc. Start()
	go http.Start()
	rpc.Start()
	log.Println("rpc srver staring")
}
