package main

import (
	"fmt"
	"golang.org/x/net/ipv4"
	"net"
)

func TraceRoute(dest string) {
	dest_addr, err := net.ResolveIPAddr("ip", dest)
	if err != nil {
		output_box.AddText(err.Error()).AddText("").RefreshText()
		stop_traceroute = true
		return
	}
	send_data := PingSendData{
		addr: dest_addr,
	}

	output_box.SetTitle(fmt.Sprintf(
		"TraceRoute to %s (%s), %d hop max",
		dest,
		dest_addr.String(),
		MaxTTL,
	))
	output_box.AddText(fmt.Sprintf(
		"TraceRoute to %s (%s), %d hop max",
		dest,
		dest_addr.String(),
		MaxTTL,
	)).RefreshText()

	for i := 1; i <= int(MaxTTL); i++ {
		if stop_traceroute {
			output_box.AddText("")
			traceroute_thread_cnt--
			return
		}

		send_data.TTL = i
		recv_data, err := Ping(send_data)
		if err != nil {
			output_box.AddText(fmt.Sprintf(
				"%d: Error, %s",
				i,
				err.Error(),
			)).RefreshText()
			continue
		}

		switch recv_data.icmp_type {
		case ipv4.ICMPTypeTimeExceeded:
			output_box.AddText(fmt.Sprintf(
				"%d: %s %.2fms",
				i,
				recv_data.name,
				recv_data.duration,
			)).RefreshText()
		case ipv4.ICMPTypeEchoReply:
			output_box.AddText(fmt.Sprintf(
				"%d: %s %.2fms",
				i,
				recv_data.name,
				recv_data.duration,
			))
			output_box.AddText(fmt.Sprintf(
				"Total hops: %d",
				i,
			)).AddText("").RefreshText()
			stop_traceroute = true
			traceroute_thread_cnt--
			return
		default:
			output_box.AddText("Unknown ICMP type").RefreshText()
		}
	}

	output_box.AddText("Too many hops").AddText("").RefreshText()
	stop_traceroute = true
	traceroute_thread_cnt--
}
