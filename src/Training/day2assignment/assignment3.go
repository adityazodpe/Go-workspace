package day2assignment

import (
	"fmt"
	"strings"
)

func Assign3() {
	fmt.Println("\nAssignment 3\n")
	fmt.Println("Enter a string:")
	var a string
	fmt.Scanln(&a)
	fmt.Println(strings.Title(strings.ToLower(a)))
}
