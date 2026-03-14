package main

import (
	"io"
	"net"
)
var partner net.Conn

func match(c net.Conn) {
	if partner != nil {
		go io.Copy(partner, c)
		go io.Copy(c, partner)
		partner = nil
	}
	partner = c
}
