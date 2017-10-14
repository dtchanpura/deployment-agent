package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dtchanpura/cd-go/manage"
)

type ipFlags []string

func (i *ipFlags) String() string {
	return ""
}

func (i *ipFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var ipList ipFlags
	addPtr := flag.Bool("add", false, "Use this to add a configuration")
	servePtr := flag.Bool("serve", false, "Use this to serve and start listener")
	repoPtr := flag.String("repopath", "", "Path to Repository")
	namePtr := flag.String("name", "", "Name of configuration")
	flag.Var(&ipList, "whitelist-ip", "Add whitelisting IPs to be added")
	serverHostPtr := flag.String("host", "0.0.0.0", "Host to bind to")
	serverPortPtr := flag.Int("port", 8080, "Port to listen on")

	flag.Parse()
	if *addPtr {
		if *namePtr == "" || *repoPtr == "" {
			log.Fatalf("Repository Name and Path required.")
		}
		manage.AddConfiguration(*namePtr, *repoPtr, "", ipList)
	} else if *servePtr {
		manage.StartServer(*serverHostPtr, *serverPortPtr)
	} else {
		fmt.Println("Atleast -add or -serve flag is required.")
	}

}
