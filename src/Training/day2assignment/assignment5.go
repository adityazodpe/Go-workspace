package day2assignment

import "fmt"

func Assign5() {
	fmt.Println("\nAssignment 5:\n")

	a5 := 20
	p := &a5

	fmt.Println("a(before)= ", a5)

	*p = 40
	fmt.Println("a(after)= ", a5)
}
