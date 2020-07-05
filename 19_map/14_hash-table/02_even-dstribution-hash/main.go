package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")

	if err != nil {
		log.Fatalln(err)
	}

	// Scan the book
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([]int, 12)

	// Loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("err: ", err)
	}

	fmt.Println(buckets)
}

func hashBucket(word string, buckets int) int {
	var sum int

	for _, v := range word {
		sum += int(v)
	}

	return sum % buckets
}

// [8517 15702 6818 8229 8172 6444 5833 11097 6687 14526 5939 9569]
