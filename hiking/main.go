package main

import "fmt"

func main() {
	steps := "UDDDUDUU"
	//var valleyTraversed int
	var altitude int
	//var mountainTraversed int
	for i := 0; i < len(steps); i++ {
		if steps[i] == 68 {
			altitude--
		} else if steps[i] == 85 {
			altitude++
		}
	}
	//fmt.Println(valleyTraversed)
	//fmt.Println(mountainTraversed)
	fmt.Println(altitude)
}
