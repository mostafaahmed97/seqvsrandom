package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func randomIO(file string) int {
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

	var chunkSize int64 = 100
	nChunks := i.Size() / chunkSize

	starts := []int64{}
	for i := range nChunks {
		starts = append(starts, i)
	}

	// shuffle slice
	for i := 0; i < len(starts); i++ {
		j := rand.Intn(i + 1)
		starts[i], starts[j] = starts[j], starts[i]
	}

	N := 0
	b := make([]byte, chunkSize)

	for _, s := range starts {
		f.Seek(s, io.SeekStart)
		n, _ := f.Read(b)
		N += n
	}

	return N
}
