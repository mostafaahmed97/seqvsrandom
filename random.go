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

	var chunkSize int64 = 1000
	nChunks := i.Size() / chunkSize

	starts := make([]int64, nChunks)
	for i := range nChunks {
		starts[i] = i
	}

	// randomize order
	rand.Shuffle(len(starts), func(i, j int) {
		starts[i], starts[j] = starts[j], starts[i]
	})

	N := 0
	b := make([]byte, chunkSize)

	for _, s := range starts {
		f.Seek(s, io.SeekStart)
		n, _ := f.Read(b)
		N += n
	}

	return N
}
