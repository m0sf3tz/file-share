package main

import "net"
import "log"
import "encoding/json"
import "fmt"
import "os"

type Message struct {
	Id   int
	Hash string
}

func Discover() {
	conn, err := net.Dial("udp", "192.168.0.255:8081")
	if err != nil {
		log.Fatal("Could not dial")
	}

	m := Message{RandId()}

	fmt.Println(m.Id)

	b, _ := json.Marshal(m)

	conn.Write(b)
}

func main() {

	conn, err := net.Dial("udp", "192.168.0.255:8081")
	if err != nil {
		log.Fatal("Could not dial")
	}

	m := Message{RandId()}

	fmt.Println(m.Id)

	b, _ := json.Marshal(m)

	conn.Write(b)

}
