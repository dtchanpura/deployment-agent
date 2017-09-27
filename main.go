package main

import (
	//"fmt"
	//"os"
	. "github.com/dtchanpura/cd-go/manage"
	"flag"
	"fmt"
	"log"
)

func main() {
	addPtr := flag.Bool("add", false, "Use this to add a configuration")
	servePtr := flag.Bool("serve", false, "Use this to serve and start listener")
	repoPtr := flag.String("repopath", "", "Path to Repository")
	namePtr := flag.String("name", "", "Name of configuration")
	serverHostPtr := flag.String("host", "0.0.0.0", "Host to bind to")
	serverPortPtr := flag.Int("port", 8080, "Port to listen on")

	flag.Parse()
	if *addPtr {
		if *namePtr == "" || *repoPtr == "" {
			log.Fatalf("Repository Name and Path required.")
		}
		AddConfiguration(*namePtr, *repoPtr, "")
	} else if *servePtr {
		StartServer(*serverHostPtr, *serverPortPtr)
	} else {
		fmt.Println("Atleast -add or -serve flag is required.")
	}

}
