package main

import "fmt"
import "os"
import "log"
import "crypto/sha256"
import "io"
import "bytes"

import "encoding/hex"

type Slice struct {
	Nth    int
	Offset int64
	Hash   []byte
	Bytes  int
}

func HeaderPrint(s []Slice) {
	for _, v := range s {
		fmt.Printf("nth    := %d \n", v.nth)
		fmt.Printf("offset := %d \n", v.offset)
		fmt.Printf("hash   := %s \n", hex.EncodeToString(v.hash[:]))
		fmt.Printf("bytes  := %d \n", v.bytes)
		fmt.Println("")
	}
}
func main() {

	slar := make([]Slice, 0)
	data := make([]byte, 1<<20) //1Meg

	file, err := os.Open("test2")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; true; i++ {

		//get the file current offset
		//must do before read since read disturbs ofset
		off, err := file.Seek(0, os.SEEK_CUR)
		if err != nil {
			log.Fatal(err)
		}
	
		//will return err=EOF.. at EOF 
		read, err := file.Read(data)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		//calculate the hash
		hash := sha256.Sum256(data[:read])

		slar = append(slar, Slice{nth: i, offset: off, hash: hash[:], bytes: read})
	}

	//HeaderPrint(slar)

	data2 := make([]byte, 1<<20) //1Meg
	f, _ := os.Open("test2")

	for i, chunk := range slar {

		f.Seek(chunk.offset, 0)
		
		br, err := f.Read(data2)
		if err == io.EOF {
			break
		}

		if br != chunk.bytes {
			panic("chunk does not match")q
		}

		h2 := sha256.Sum256(data2[:br])

		if !bytes.Equal(h2[:], chunk.hash) {
			panic(0)
		}

		fmt.Printf("nth    := %d \n", chunk.nth)
		fmt.Println("calculated")
		fmt.Printf("offset := %d \n", chunk.offset)
		fmt.Printf("hash   := %s \n", hex.EncodeToString(chunk.hash[:]))
		fmt.Println("Derived")
		fmt.Printf("hash   := %s \n", hex.EncodeToString(h2[:]))
		fmt.Println("")

	}
}
