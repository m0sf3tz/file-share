package main

import "fmt"
import "encoding/json"

type Slice struct {
	Nth   int
	Bytes int
}

func main() {
	foo := make([]Slice, 0)
	foo = append(foo, Slice{20, 30})
	foo = append(foo, Slice{2, 3430})

	jsonData, err := json.Marshal(foo)
	fmt.Println(err)
	fmt.Println(jsonData)

	bar := make([]Slice, 0)
	json.Unmarshal(jsonData, &bar)
	fmt.Println(bar)
}
