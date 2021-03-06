package day2assignment

import (
	"fmt"
	"strconv"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func Assign2() {
	fmt.Println("\nAssignment 2\n")
	fmt.Println("Give an input:")
	var input string
	fmt.Scanln(&input)
	if isNumeric(input) == false {
		fmt.Printf("Before converting, value = %v, type = %T\n", input, input)
		fmt.Println("No need for conversion.")
	} else {
		i1, err := strconv.Atoi(input)
		if err == nil {
			fmt.Printf("Before converting, value = %v, type = %T\n", input, input)
			fmt.Printf("After converting, value = %v,type = %T\n", i1, i1)
		} else {
			s, err1 := strconv.ParseFloat(input, 64)
			if err1 == nil {
				fmt.Printf("Before converting, value = %v, type = %T\n", input, input)
				fmt.Printf("After converting, value = %v,type = %T\n", s, s)
			}
		}
	}
}
