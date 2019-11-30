package main

import "fmt"
import "encoding/json"

type WN_s struct {
	FileArr []File_s
}

type File_s struct {
	Name string
	Hash int
}

func main() {

	type JsonMaster struct {
		Needs WN_s
		Wants WN_s
	}

	j := JsonMaster{
		Needs: WN_s{
			FileArr: []File_s{
				{Name: "FilaA", Hash: 200},
				{"FilaB", 450},
			},
		},
		Wants: WN_s{
			FileArr: []File_s{
				{"FilaC", 500},
			},
		},
	}

	var jsonData []byte

	jsonData, _ = json.MarshalIndent(j, "", "    ")

	fmt.Println(string(jsonData))
}
