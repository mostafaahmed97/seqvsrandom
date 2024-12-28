# Sequential vs Random Disk I/O Comparison

> Free Palestine, End the occupation 
> 
> [![Ceasefire Now](https://badge.techforpalestine.org/default)](https://techforpalestine.org/learn-more)

A simple Go experiment to measure and compare the performance difference between sequential and random disk I/O operations.

Read the [blog post](https://mostafaahmed.hashnode.dev/is-random-disk-io-really-that-bad)

# Usage

After running `go build`, the following commands are available

## Generate Test Data

```go
seqvsrandom gen <file> <length>
```

Parameters:
- `file`: Path to output file
- `length`: File size in bytes (optional, defaults to 10K bytes)

## Sequential Read Test

```go
seqvsrandom seq <file>
```
Parameters:
- `file`: Path to input file

Reads the entire file sequentially and reports the time taken.

## Random Read Test

```go
seqvsrandom rnd <file>
```
Parameters:
- `file`: Path to input file

Reads the file in random 1000-byte chunks and reports the time taken.

# Building

```bash
go build
```

# Example

```bash
# Generate a 10MB test file
./seqvsrandom gen data.txt 10000000

# Run sequential read test
./seqvsrandom seq data.txt

# Run random read test
./seqvsrandom rnd data.txt
```
![alt text](usage.png)

