package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const input = "Laziness is the mother of all bad habits, but ultimately she is a mother and we should respect her ― Shikamaru Nara"

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(os.Stderr, "reading input:", err)
	}
}

// Laziness
// is
// the
// mother
// of
// all
// bad
// habits,
// but
// ultimately
// she
// is
// a
// mother
// and
// we
// should
// respect
// her
// ―
// Shikamaru
// Nara
