package main

import "fmt"

func main() {
	birds := []int32{1, 4, 4, 4, 5, 3, 5, 5}
	freqArr := [6]int32{}

	for i := 0; i < len(birds); i++ {
		freqArr[birds[i]]++
	}
	//fmt.Println(freqArr)
	for j := 0; j < (len(freqArr)); j++ {
		for k := 1; k < len(freqArr); k++ {
			if freqArr[j] < freqArr[k] {
				freqArr[j] = 0
			}

		}
	}
	//fmt.Println(freqArr)
	//var highFreqBird int32
	for l := 0; l < len(freqArr); l++ {
		if freqArr[l] > 0 {
			fmt.Println(l)
			break
		}
	}
}
