package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

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

	fmt.Println(buckets)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])

	bucket := letter % buckets

	return bucket
}

// [7909 34850 14167 22516 11229 6585 13167 20653 47306 14912 4503 16315]
