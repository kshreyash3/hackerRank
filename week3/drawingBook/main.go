package main

import "fmt"

func main() {
	n := 10
	p := 6
	flip := p / 2
	flip1 := (n - p) / 2
	if flip1 > flip {
		fmt.Println(flip)
	} else {
		fmt.Println(flip1)
	}

}
