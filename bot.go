package main

import (
	"io"
	"time"

	"go.sakib.dev/chat/markov"
)

type bot struct {
	io.ReadCloser
	out io.Writer
}

func (b bot) Write(buf []byte) (int, error) {
	go b.speak()
	return len(buf), nil
}

func Bot() io.ReadWriteCloser {
	r, out := io.Pipe()
	return bot{r, out}
}

var chain = markov.NewChain(2)

func (b bot) speak() {
	time.Sleep(time.Second)
	msg := chain.Generate(10)
	b.out.Write([]byte(msg + "\n"))
}