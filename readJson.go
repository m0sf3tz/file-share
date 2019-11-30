package main

import "io/ioutil"
import "encoding/json"
import "fmt"

type JsonMaster struct {
	Needs WN_s
	Wants WN_s
}

type WN_s struct {
	FileArr []File_s
}

type File_s struct {
	Name string
	Hash int
}

func main() {
	b, err := ioutil.ReadFile("j.json")
	if err != nil {
		panic(0)
	}

	var j JsonMaster
	json.Unmarshal(b, &j)

	fmt.Println(j)

	for _, v := range j.Needs.FileArr {
		fmt.Println(v)
	}
}
