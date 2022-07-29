package main

import (
	"fmt"
	"golang.org/x/net/ipv4"
	"net"
	"time"
)

var (
	sec = time.Second
	// from https://pkg.go.dev/golang.org/x/net/internal/iana
	ProtocolICMP      = 1
	MaxTTL       int8 = 30
)

func main() {
	dest := "google.com"
	dest_addr, err := net.ResolveIPAddr("ip", dest)
	if err != nil {
		panic(err)
	}
	send_data := PingSendData{
		addr: dest_addr,
	}

traceroute_loop:
	for i := 1; i < int(MaxTTL); i++ {
		send_data.seq = i
		recv_data := Ping(send_data)
		switch recv_data.icmp_type {
		case ipv4.ICMPTypeTimeExceeded:
			fmt.Printf("%d: %s %.2fms\n", i, recv_data.name, recv_data.duration)
		case ipv4.ICMPTypeEchoReply:
			fmt.Printf("%d: %s %.2fms\n", i, recv_data.name, recv_data.duration)
			break traceroute_loop
		default:
			fmt.Printf("Unknown ICMP type\n")
		}
		time.Sleep(1 * sec)
	}
}
