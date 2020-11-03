package main

import "net"

func main() {
	conn, _ := net.Dial("udp", "192.168.10.1:8889")
	conn.Write([]byte("command"))
	conn.Write([]byte("takeoff"))
}
