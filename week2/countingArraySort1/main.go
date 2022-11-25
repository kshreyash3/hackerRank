package main

import "fmt"

func main() {
	arr := []int32{1, 1, 3, 2, 1}
	freqArr := [100]int32{}
	for i := 0; i < len(arr); i++ {
		freqArr[arr[i]]++
	}
	fmt.Println(freqArr)
}
