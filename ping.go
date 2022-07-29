package main

import (
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

type PingSendData struct {
	addr *net.IPAddr
	seq  int
}

type PingRecvData struct {
	icmp_type icmp.Type
	duration  float64
	name      string
}

func Ping(ping_ctrl PingSendData) PingRecvData {
	var (
		send_msg = icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				Seq: ping_ctrl.seq,
				ID:  os.Getpid() & 0xFFFF,
			},
		}
		read_buffer = make([]byte, 1024)
		duration    float64
		name        string
	)

	// conn_tmp is net.PacketConn
	// conn is *net.PacketConn (Using c as its underlying transport)
	conn_tmp, _ := net.ListenPacket("ip4:icmp", "0.0.0.0")
	defer conn_tmp.Close()
	conn := ipv4.NewPacketConn(conn_tmp)

	write_buffer, err := send_msg.Marshal(nil)
	if err != nil {
		panic(err)
	}
	conn.SetTTL(ping_ctrl.seq)

	startTime := time.Now()
	conn.WriteTo(write_buffer, nil, ping_ctrl.addr)
	conn.SetDeadline(time.Now().Add(3 * sec))
	n, _, src, err := conn.ReadFrom(read_buffer)
	if err != nil {
		panic(err)
	}
	duration = float64(time.Since(startTime).Nanoseconds()) / 1000000

	read_msg, err := icmp.ParseMessage(ProtocolICMP, read_buffer[:n])

	names, err := net.LookupAddr(src.String())
	if err != nil {
		name = src.String()
	} else {
		name = names[0]
	}

	return PingRecvData{
		icmp_type: read_msg.Type,
		duration:  duration,
		name:      name,
	}
}
