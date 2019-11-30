package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 3 {
			fmt.Println("here")

			fmt.Println(&e)
			l.Remove(e)
			fmt.Println(&e)
			fmt.Println(e.Value)
		}
	}

	fmt.Println("")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
