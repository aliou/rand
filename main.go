package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomSequence(length int) string {
	sequence := make([]rune, length)

	for i := range sequence {
		sequence[i] = letters[rand.Intn(len(letters))]
	}

	return string(sequence)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randomSequence(5))
}
