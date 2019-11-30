package main

import "net"
import "log"
import "encoding/json"
import "fmt"
import "os"

func Discover() {
	fmt.Println("Starting Discover seqeunce..")

	conn, err := net.Dial("udp", "192.168.0.255:8081")
	if err != nil {
		log.Fatal("Could not dial")
	}

	m := Message{RandId(), os.Args[1]}

	fmt.Println(m.Id)

	b, _ := json.Marshal(m)

	conn.Write(b)
}

func main() {

	Discover()

}
