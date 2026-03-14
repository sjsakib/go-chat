package main

import (
	"log"
	"net"
)

const listenAddr = "localhost:4000"

func runTcp() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		match(c)
	}
}
