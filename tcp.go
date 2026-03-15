package main

import (
	"io"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

type tcpConn struct {
	io.WriteCloser
	io.Reader
}

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
		r := feedChain(c)
		go match(tcpConn{c, r})
	}
}

func feedChain(rw io.ReadWriter) io.Reader {
	r, w := io.Pipe()
	go func() {
		_, err := io.Copy(io.MultiWriter(w, chain), rw)
		if err != nil {
			w.CloseWithError(err)
		}
	}()
	return r
}
