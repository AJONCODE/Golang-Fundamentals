package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)

	defer res.Body.Close()

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 200)

	// loop over the words
	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}

	for i := 65; i < 126; i++ {
		fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	}
}

func hashBucket(word string) int {
	return int(word[0])
}

// 65 - A - 1790
// 66 - B - 1260
// 67 - C - 910
// 68 - D - 522
// 69 - E - 325
// 70 - F - 658
// 71 - G - 473
// 72 - H - 812
// 73 - I - 2829
// 74 - J - 244
// 75 - K - 80
// 76 - L - 504
// 77 - M - 563
// 78 - N - 672
// 79 - O - 433
// 80 - P - 816
// 81 - Q - 291
// 82 - R - 259
// 83 - S - 1671
// 84 - T - 1771
// 85 - U - 87
// 86 - V - 119
// 87 - W - 1080
// 88 - X - 5
// 89 - Y - 205
// 90 - Z - 15
// 91 - [ - 21
// 92 - \ - 0
// 93 - ] - 0
// 94 - ^ - 0
// 95 - _ - 2
// 96 - ` - 0
// 97 - a - 21787
// 98 - b - 9692
// 99 - c - 7385
// 100 - d - 5219
// 101 - e - 3614
// 102 - f - 7576
// 103 - g - 3016
// 104 - h - 12620
// 105 - i - 11607
// 106 - j - 582
// 107 - k - 852
// 108 - l - 5321
// 109 - m - 7765
// 110 - n - 4060
// 111 - o - 13798
// 112 - p - 5256
// 113 - q - 395
// 114 - r - 3592
// 115 - s - 16260
// 116 - t - 33335
// 117 - u - 2591
// 118 - v - 1452
// 119 - w - 13314
// 120 - x - 1
// 121 - y - 2272
// 122 - z - 19
// 123 - { - 0
// 124 - | - 0
// 125 - } - 0
