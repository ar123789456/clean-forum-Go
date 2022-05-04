package main

import (
	"forum/client"
	"forum/server"
	"log"
)

func main() {
	go server.Run()
	client.Run()
	log.Println("serv")
}
