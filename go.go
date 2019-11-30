package main

import "fmt"
import "time"

func test(p <-chan string) {
	s := <-p
	fmt.Printf("got a %s \n", s)
}

func main() {
	q := make(chan string)

	go test(q)

	q <- "what"
	time.Sleep(time.Duration(10) * time.Second)
}
