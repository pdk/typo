package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const spaces = "          "
const digits = "0123456789"
const punc = "\\\"'~@#$%^&*()!=_+-|':;`'{}[]<>?./,"
const letters0 = "abcdefghijklmnopqrstuvwxyz"
const letters1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const pickings = spaces + punc + letters0 + letters0 + letters0 + letters1

// const pickings = spaces + digits + punc + letters0 + letters0 + letters0 + letters1
// const pickings = spaces + letters0 + letters0 + letters0 + letters1
// const pickings = digits + digits + digits + digits + digits + digits + digits + "*+/-=."

func randSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = pickings[rand.Intn(len(pickings))]
	}
	return string(b)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		sample := randSeq(75)
		fmt.Println(sample)
		text, _ := reader.ReadString('\n')

		if text == "" {
			break
		}

		good, bad := 0, 0

		for i := range sample {
			if i >= len(text) {
				break
			}

			if sample[i] == text[i] {
				fmt.Print(" ")
				good = good + 1
			} else {
				fmt.Print("^")
				bad = bad + 1
			}
		}

		fmt.Println("  -- ", good, " / ", bad)
	}
}
