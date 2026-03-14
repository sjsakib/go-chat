package main

import (
	"fmt"
	"io"
)
var partner io.ReadWriteCloser

func match(c io.ReadWriteCloser) {
	errCh := make(chan error, 1)
	if partner != nil {
		cPartner := partner
		partner = nil
		
		go copy(cPartner, c, errCh)
		go copy(c, cPartner, errCh)

		if err := <-errCh; err != nil {
			fmt.Println(err)
		}
		cPartner.Close()
		c.Close()

		return
	}
	partner = c
}
