package main

import "fmt"

func main() {
	sticks := []int32{1, 2, 3, 4, 5, 10}
	triangleArr := [4]int32{}

	for i := 0; i < len(sticks); i++ {
		for j := i + 1; j < (len(sticks)); j++ {
			for k := j + 1; k < len(sticks); k++ {
				if sticks[i]+sticks[j] > sticks[k] && sticks[j]+sticks[k] > sticks[i] && sticks[i]+sticks[k] > sticks[j] {
					//fmt.Println(sticks[i], sticks[j], sticks[k])
					triangleArr[0] = sticks[i]
					triangleArr[1] = sticks[j]
					triangleArr[2] = sticks[k]
					triangleArr[3] = sticks[i] + sticks[j] + sticks[k]
					fmt.Println(triangleArr)

					// perimeterSum := triangleArr[0] + triangleArr[1] + triangleArr[2]
					// fmt.Println(perimeterSum)

				}
			}

		}

	}
}
