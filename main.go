package main

import (
	"github.com/yakovasavr/sql_connection/webserver"
)

func main() {
	var server webserver.WEBServer
	server.Start()
}
