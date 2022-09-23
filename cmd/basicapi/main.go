package main

import (
	"flag"
	"go-api/cmd/server"
)

func main() {
	httpPort := flag.String("http_port", "8080", "the port for the server")
	server.RunServer(*httpPort)
}
