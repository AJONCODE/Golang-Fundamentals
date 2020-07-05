package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice of slice of string to hold slices of words
	buckets := make([][]string, 12)

	fmt.Println(buckets)

	// code here has been updated from the recording
	// see below for explanation

	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// Print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}

	// fmt.Println(buckets)

	fmt.Println(len(buckets))
	fmt.Println(cap(buckets))
}

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}

// [[] [] [] [] [] [] [] [] [] [] [] []]
// 0  -  8517
// 1  -  15702
// 2  -  6818
// 3  -  8229
// 4  -  8172
// 5  -  6444
// 6  -  5833
// 7  -  11097
// 8  -  6687
// 9  -  14526
// 10  -  5939
// 11  -  9569
// 12
// 12
