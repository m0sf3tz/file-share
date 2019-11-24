package main

import "math/rand"
import "time"

func RandId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000000)
}
