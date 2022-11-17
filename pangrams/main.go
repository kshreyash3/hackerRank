package main

import (
	"fmt"
	"strings"
)

var a bool

func main() {
	var s string
	alphabets := "abcdfeghijklmnopqrstuvwxyz"
	arr := [26]int{}
	input := "We promptly judged antique ivory buckles for the next prize"
	input = strings.ToLower(input)

	for i := 0; i < len(alphabets); i++ {
		for j := 0; j < len(input); j++ {
			if alphabets[i] == input[j] {
				arr[i]++
			}
		}
	}

	for k := 0; k < len(arr); k++ {
		if arr[k] == 0 {
			s = ("not pangram")
		}
	}
	if s != ("not pangram") {
		s = "pangram"
	}
	fmt.Println(s)

}
