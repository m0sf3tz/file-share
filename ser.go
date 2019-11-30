package main

import "fmt"
import "net"
import "log"
import "encoding/json"

func main() {
	ln, err := net.ListenPacket("udp", ":8081")
	if err != nil {
		log.Fatal("Could not start listening")
	}

	fmt.Println(ln.LocalAddr())

	b := make([]byte, 256)

	for {

		n, addr, err := ln.ReadFrom(b)
		if err != nil {
			log.Fatal("Could accept")
		}

		fmt.Printf("Incomming from %s \n ", addr.String())

		var m Message

		var er error = json.Unmarshal(b[0:n], &m)
		if er != nil {
			fmt.Println("shjit")
		}

		fmt.Println(m)
	}

}
