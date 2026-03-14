package main

import "golang.org/x/net/websocket"

type socket struct {
	conn *websocket.Conn
	done chan bool
}

func (s socket) Read(b []byte) (int, error)  { return s.conn.Read(b) }
func (s socket) Write(b []byte) (int, error) { return s.conn.Write(b) }

func (s socket) Close() error {
	s.done <- true
	return nil
}

func socketHandler(ws *websocket.Conn) {
	s := socket{conn: ws, done: make(chan bool)}
	go match(s)
	<-s.done
}
