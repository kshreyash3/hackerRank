package main

import (
	"fmt"
	"strings"
)

func main() {
	alphabets := "abcdfeghijklmnopqrstuvwxyz"
	input := "We promptly judged antique ivory buckles for the prize"
	input = strings.ToLower(input)

	var counter int
	fmt.Println(input)
	for i := 0; i < len(alphabets); i++ {
		for j := 0; j < len(input); j++ {
			if alphabets[i] == input[j] {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
