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
	r, w := io.Pipe()
	go func() {
		_, err := io.Copy(io.MultiWriter(w, chain), ws)
		if err != nil {
			w.CloseWithError(err)
		}
	}()
	s := socket{r, ws, make(chan bool)}
	go match(s)
	<-s.done
}
