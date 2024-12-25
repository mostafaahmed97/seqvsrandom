package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tseqvsrandom gen <file> <length>")
	fmt.Println("\t\t- <length>: integer, generated file length in bytes, defaults to 1M bytes (1MB)")
	fmt.Println("\t\t- <file>: path to file")
	fmt.Println("\tseqvsrandom seq <file>")
	fmt.Println("\tseqvsrandom rnd <file>")
}

func totalTime(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("total time -> %s: %s\n", name, elapsed)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	mode := os.Args[1]
	file := os.Args[2]

	if mode == "gen" {
		fsize := 1_000_000
		if len(os.Args) > 3 {
			fsize, _ = strconv.Atoi(os.Args[3])
		}

		fmt.Printf("writing %d bytes to %s\n", fsize, file)
		n := generate(file, fsize)
		fmt.Printf("done, %d bytes written\n", n)

	} else if mode == "seq" {
		defer totalTime(time.Now(), "SequentialIO")

		n := sequentialIO(file)
		fmt.Printf("read %d bytes sequentially\n", n)
	} else if mode == "rnd" {
		defer totalTime(time.Now(), "RadomIO")

		n := randomIO(file)
		fmt.Printf("read %d bytes randomly\n", n)
	} else {
		printUsage()
	}
}
