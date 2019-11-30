package main

import "fmt"
import "os"
import "log"
import "crypto/sha256"
import "io"

import "encoding/hex"

type slice struct {
	nth    int
	offset int64
	hash   []byte
	bytes  int
}

func main() {

	slar := make([]slice, 0)
	data := make([]byte, 2<<20) //1Meg

	file, err := os.Open("test2")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; true; i++ {
		read, err := file.Read(data)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if read == 0 {
			break
		}

		//get the file current offset
		off, err := file.Seek(0, os.SEEK_CUR)
		if err != nil {
			log.Fatal(err)
		}

		//calculate the hash
		hash := sha256.Sum256(data[:read])

		slar = append(slar, slice{nth: i, offset: off, hash: hash[:], bytes: read})
	}

	for _, v := range slar {
		fmt.Printf("nth    := %d \n", v.nth)
		fmt.Printf("offset := %d \n", v.offset)
		fmt.Printf("hash   := %s \n", hex.EncodeToString(v.hash[:]))
		fmt.Println("")
	}
}

/*
	_, err = file.Read(data)
	hash = sha256.Sum256(data)
	fmt.Println(hex.EncodeToString((hash[:])))
*/
