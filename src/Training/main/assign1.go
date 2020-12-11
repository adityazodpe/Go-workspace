package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Assignment 1:\n")
	a := 14.65478
	fmt.Printf("%.2f\n", a)

	fmt.Println("\nAssignment 2:\n")
	input := "32"
	n, err := strconv.Atoi(input)
	if err == nil {
		fmt.Printf("Before converting. Value=%v, Type=%T\n", input, input)
		fmt.Printf("After converting. Value=%v, Type=%T\n\n", n, n)
	}

	input2 := "32.56"
	n2, _ := strconv.ParseFloat(input2, 64)
	fmt.Printf("Before converting. Value=%v, Type=%T\n", input2, input2)
	fmt.Printf("After converting. Value=%v, Type=%T\n\n", n2, n2)

	fmt.Println("Assignment 3:\n")
	input3 := "go-lang-training"
	output := strings.Title(input3)
	fmt.Println("Input: ", input3)
	fmt.Println("Output: ", output)

	fmt.Println("\nAssignment 4:\n")
	const (
		x = 6
		y
		z
	)
	fmt.Println("x=", x, "y=", y, "z=", z)

	fmt.Println("\nAssignment 5:\n")

	a5 := 20
	p := &a5

	fmt.Println("a(before)= ", a5)

	*p = 40
	fmt.Println("a(after)= ", a5)

	fmt.Println("\nAssignment 6:\n")
	str1 := "1Ax3 7y Bk"
	res := strings.Count(str1, " ")
	fmt.Println("No. of spaces: ", res)
}
