package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	ttyS, err := os.OpenFile("/dev/ttyS0", os.O_RDWR, 0666)
	if err != nil {
		log.Print(err)
		return
	}

	go func() {
		inbox := make([]byte, 64)
		for {
			n, err := ttyS.Read(inbox)
			fmt.Print(inbox[:n])
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
	}()

	stdin := make([]byte, 64)
	for {
		n, err := os.Stdin.Read(stdin)
		if err == io.EOF {
			os.Exit(0)
		} else if err != nil {
			log.Fatal(err)
		}

		_, err = ttyS.Write(stdin[:n])
		if err != nil {
			log.Fatal(err)
		}
	}
}
