package main

import "fmt"

func main() {
	var n, i, pair, q int32
	n = 10
	ar := [10]int32{1, 2, 1, 2, 1, 3, 2, 2, 3, 4}
	freqArray := [len(ar)]int32{}

	for i = 0; i < n; i++ {
		freqArray[ar[i]]++
	}
	fmt.Println(freqArray)
	for j := 0; j < len(freqArray); j++ {
		if freqArray[j]%2 == 1 {
			freqArray[j]--
		}
	}
	fmt.Println(freqArray)
	for k := 0; k < len(freqArray); k++ {
		if freqArray[k] > 0 {
			if freqArray[k] == 2 {
				pair++
			} else if freqArray[k] > 3 {
				q = freqArray[k] / 2
				pair += q
			}
		}

	}
	fmt.Println(pair)
}
