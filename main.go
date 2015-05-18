package main

import (
	"./cron"
	"./http"
	"./rpc"
	//"flag"
	//"fmt"
	// "github.com/dinp/server/g"
	// "github.com/dinp/server/hbs"
	// "github.com/dinp/server/http"
	//"os"
	"log"
)

func main() {
	/*cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}*/

	// g.ParseConfig(*cfg)

	// g.InitRedisConnPool()
	// g.InitDbConnPool()
	log.Println("staring...")
	go cala.Timeshow()

	go http.Start()
	rpc.Start()
	log.Println("rpc srver staring")
}
