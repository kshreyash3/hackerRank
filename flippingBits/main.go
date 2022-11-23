package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	number := 3113
	binNumber := fmt.Sprintf("%b", number)
	fmt.Println(binNumber)
	mylen := 32 - len(binNumber)
	fmt.Println(mylen)
	binNumberArr := make([]string, len(binNumber), 32)
	appendArr := make([]string, mylen, 32)
	//binNumberArr := []string{}
	//appendArr := []string{}
	for i := 0; i < len(binNumber); i++ {
		binNumberArr[i] = string(binNumber[i])
	}
	fmt.Println(binNumberArr)

	for k := 0; k < len(binNumberArr); k++ {
		if binNumberArr[k] == "0" {
			binNumberArr[k] = "1"
		} else {
			binNumberArr[k] = "0"
		}
	}

	for j := 0; j < mylen; j++ {
		appendArr[j] = "1"
	}
	fmt.Println(appendArr)
	binNumberArr = append(appendArr, binNumberArr...)
	binNumber = fmt.Sprintf(strings.Join(binNumberArr, ""))
	fmt.Println(binNumber)

	if final, err := strconv.ParseInt(binNumber, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(final)
	}
}
