package main

import "fmt"

func main() {
	sticks := []int32{1, 2, 3, 4, 5, 10}
	fmt.Println(possibleTriangle(sticks))
}

func possibleTriangle(arr []int32) []int32 {

	triangleArr := [3]int32{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < (len(arr)); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j] > arr[k] && arr[j]+arr[k] > arr[i] && arr[i]+arr[k] > arr[j] {
					//fmt.Println(arr[i], arr[j], arr[k])
					triangleArr[0] = arr[i]
					triangleArr[1] = arr[j]
					triangleArr[2] = arr[k]
				}
			}

		}

	}
	return triangleArr[:]
}
