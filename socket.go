package main

import (
	"io"

	"golang.org/x/net/websocket"
)

type socket struct {
	io.Reader
	io.Writer
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func socketHandler(ws *websocket.Conn) {
	r := feedChain(ws)
	s := socket{r, ws, make(chan bool)}
	go match(s)
	<-s.done
}
