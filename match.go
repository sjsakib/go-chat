package main

import (
	"fmt"
	"io"
	"time"
)

var partnerCh = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	select {
	case partner := <-partnerCh:
		chat(c, partner)
	case partnerCh <- c:
	case <-time.After(5 * time.Second):
		chat(Bot(), c)
	}

}

func chat(c1, c2 io.ReadWriteCloser) {
	fmt.Fprintf(c1, "Found match!")
	fmt.Fprintf(c2, "Found match!")
	errCh := make(chan error, 1)
	go copy(c1, c2, errCh)
	go copy(c2, c1, errCh)

	if err := <-errCh; err != nil {
		fmt.Println(err)
	}
	c1.Close()
	c2.Close()
}
