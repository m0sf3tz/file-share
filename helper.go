package main

import "math/rand"
import "time"

type Message struct {
	Id   int
	Hash string
}

func RandId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000000)
}
