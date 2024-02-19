package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for i := 0; i < 10000; i++ {
		file, err := os.Create(fmt.Sprintf("./tmp/file_%d.txt", i))
		if err != nil {
			log.Panic(err)
		}
		defer file.Close()

		file.WriteString("Hello, World!")
	}
}
