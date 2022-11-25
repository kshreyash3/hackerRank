package main

import "fmt"

func main() {
	arr := [7]int32{7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}

	}
	fmt.Println(arr)
	fmt.Println(arr[len(arr)/2])
}
