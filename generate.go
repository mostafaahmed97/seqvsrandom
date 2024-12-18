package main

import (
	"fmt"
	"os"
)

func generate(file string, fsize int) int {
	n := 0

	f, err := os.Create(file)
	if err != nil {
		fmt.Printf("failed to open file %s\n", file)
		panic(err)
	}

	defer f.Close()

	for i := 1; i <= fsize; i++ {
		if i%10 == 0 {
			f.WriteString("\n")
		} else {
			f.WriteString("A")
		}

		n++
	}

	err = f.Sync()
	if err != nil {
		fmt.Printf("failed to flush file to disk\n")
		panic(err)
	}

	return n
}
