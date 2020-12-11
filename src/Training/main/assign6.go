package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "1Ax3 7y Bk"
	res := strings.Count(str1, " ")

	fmt.Println("number of spaces are ", res)
}
