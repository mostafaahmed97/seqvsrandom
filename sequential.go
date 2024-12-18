package main

import (
	"fmt"
	"io"
	"os"
)

func sequentialIO(file string) int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("failed to read file in sequential mode")
		panic(err)
	}

	defer f.Close()

	i, err := f.Stat()
	if err != nil {
		fmt.Println("failed to read file info")
		panic(err)
	}

	f.Seek(0, io.SeekStart)

	b := make([]byte, i.Size())

	n, err := f.Read(b)
	if err != nil {
		fmt.Println("failed to read file contents")
		panic(err)
	}

	return n
}
