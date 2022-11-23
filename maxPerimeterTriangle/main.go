package main

import "fmt"

func main() {
	sticks := []int32{1, 2, 3, 4, 5, 10}
	possibleTriangle(sticks)
}

func possibleTriangle(arr []int32) triangle {
	type triangle struct {
		l1 int32
		l2 int32
		l3 int32
	}
	var t triangle
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < (len(arr)); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j] > arr[k] && arr[j]+arr[k] > arr[i] && arr[i]+arr[k] > arr[j] {
					//fmt.Println(arr[i], arr[j], arr[k])
					t := triangle{arr[i], arr[j], arr[k]}
				} else {
					fmt.Println(-1)
				}
			}

		}

	}
	return t
}
