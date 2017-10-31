package main

import (
	"log"
	"github.com/werpas/go-template-svc/rest"
)

func main() {
	log.Println("Starting Go Service ...")

	//  start rest interface here
	rest.StartServer()
}