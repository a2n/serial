package main

import (
	"github.com/a2n/serial/src"
	"github.com/a2n/serial/src/grpc/server"
)

func main() {
	web := serial.NewWebService()
	grpc := server.NewSerialServer()

	ch := make(chan struct{})
	go web.Start()
	go grpc.Start()
	<-ch
}
