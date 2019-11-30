package main

import "fmt"

func main() {
	var s []string

	s = append(s, "cock")
	s = append(s, "sucker!")
	p(s)
}

func p(s []string) {
	for _, val := range s {
		fmt.Printf("the val is %s \n", val)
	}
}
