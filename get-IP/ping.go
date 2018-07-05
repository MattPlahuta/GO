package main

import (
	"net"
	"fmt"
)

func main() {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    // handle err...
	_ = err
	
     defer conn.Close()
     localAddr := conn.LocalAddr().(*net.UDPAddr)
	 
	 fmt.Println(conn)
	 fmt.Println(localAddr)
}