package main

import (
	"log"
	"net"
)

func main() {
	s := newsever()
	go s.run()

	listener,err := net.Listne("tcp",":8888")
	if err !=nil{
		log.Fatalf("unable to start server: %s",err.Error())
	}

	defer listener.Close()
	log.Printf("sterted server on :8888")

	for{
		conn,err := listener.Accept()
		if err != nil{
			log/Print("unable to accept connection: %s",err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
