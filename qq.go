package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sleeperz(c chan int) {
	s := rand.Intn(10)

	fmt.Printf("Sleeping for %d seconds \n", s)
	time.Sleep(time.Duration(s) * time.Second)
	c <- 1
}

func test() {

	rand.Seed(time.Now().UnixNano())

	res := make(chan int)
	go sleeperz(res)
	go sleeperz(res)
	<-res
}

func main() {
	test()

	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
