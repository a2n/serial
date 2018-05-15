package main

import (
	"flag"

	"github.com/a2n/serial/src"
	"github.com/a2n/serial/src/grpc/server"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":57888", "grpc port")
}

func main() {
	flag.Parse()

	web := serial.NewWebService()
	grpc := server.NewSerialServer()

	ch := make(chan struct{})
	go web.Start()
	go grpc.Start(port)
	<-ch
}
