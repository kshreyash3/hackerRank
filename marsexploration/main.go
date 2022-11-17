package main

import "fmt"

func main() {
	receivedSignal := "SOSSPSSQSSOR"
	var deviation int32
	Times := (len(receivedSignal) / 3)
	//fmt.Println(Times)
	expectedSignal := "SOS"
	for i := 0; i < Times-1; i++ {
		expectedSignal += "SOS"
	}
	//fmt.Println(expectedSignal)
	for j := 0; j < len(receivedSignal); j++ {
		if receivedSignal[j] != expectedSignal[j] {
			deviation++
		}
	}
	fmt.Println(deviation)
}
