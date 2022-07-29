package main

import (
	"golang.org/x/net/ipv4"
	"fmt"
	"net"
)

func TraceRoute(dest string) {
	dest_addr, err := net.ResolveIPAddr("ip", dest)
	if err != nil {
		output_box.AddText(err.Error())
		output_box.RefreshText()
		app.Draw()
	}
	send_data := PingSendData{
		addr: dest_addr,
	}

	for i := 1; i <= int(MaxTTL); i++ {
		if stop_traceroute {
			return
		}

		send_data.seq = i
		recv_data := Ping(send_data)

		switch recv_data.icmp_type {
		case ipv4.ICMPTypeTimeExceeded:
			output_box.AddText(fmt.Sprintf(
				"%d: %s %.2fms",
				i,
				recv_data.name,
				recv_data.duration))
			output_box.RefreshText()
		case ipv4.ICMPTypeEchoReply:
			output_box.AddText(fmt.Sprintf(
				"%d: %s %.2fms",
				i,
				recv_data.name,
				recv_data.duration))
			output_box.AddText(fmt.Sprintf(
				"Total hops: %d",
				i))
			output_box.RefreshText()
			return
		default:
			output_box.AddText("Unknown ICMP type")
			output_box.RefreshText()
		}
	}
}
