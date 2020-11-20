package main

/*
  Rules:
  Write your solution using Go Programming Language
  Your source code must be a single file (package main)
  Do not use any <for> statement
  You may only use standard library packages
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getUint(scanner *bufio.Scanner) uint64 {
	scanner.Scan()
	num, _ := strconv.ParseUint(scanner.Text(), 10, 32)
	return num
}

func iterateThrough(n uint64, cb func(uint64)) {
	var counter uint64 = 0
startIteration:
	if counter == n {
		return
	}
	cb(counter)
	counter++
	goto startIteration
}

func calcSquareSum(scanner *bufio.Scanner, n uint64) uint64 {
	scanner.Scan()
	tokens := strings.SplitN(scanner.Text(), " ", int(n+1))
	var sum uint64 = 0
	iterateThrough(n, func(index uint64) {
		num, _ := strconv.ParseUint(tokens[index], 10, 32)
		sum += uint64(math.Pow(float64(num), 2))
		//fmt.Printf("Sum: %v", sum)
	})
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	N := getUint(scanner)
	var results []uint64
	iterateThrough(N, func(index uint64) {
		X := getUint(scanner)
		results = append(results, calcSquareSum(scanner, X))
	})
	iterateThrough(N, func(index uint64) {
		fmt.Println(results[index])
	})
}
