package main

import "fmt"
import "os"
import "strconv"
import "log"

//editNth fileName, offset
func main() {

	fmt.Println(len(os.Args))

	if len(os.Args) != 3 {
		fmt.Println("need to provide EXACTLY three argument")
	}

	i, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("must provide string as argument")
	}

	fmt.Println(i)

	fmt.Printf("opening.. %s \n", os.Args[1])
	file, err := os.OpenFile(os.Args[1], os.O_RDWR, os.ModeAppend) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The size of the file is %d \n", info.Size())

	ar := []byte{1, 2, 3}
	_, err = file.WriteAt(ar, int64(i))
	fmt.Println(err)
	file.Sync()

}
