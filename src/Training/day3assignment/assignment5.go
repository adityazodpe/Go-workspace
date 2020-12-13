package day3assignment

import "fmt"

func Assign5() {
	fmt.Println("\nAssignment 5:")
	num := [7]int{2, 3, 5, 7, 9, 11, 34}
	var s []int = num[0:2]
	var s2 []int = num[4:6]

	s3 := append(s, s2...)

	fmt.Println(num)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println("Concatenating 2 slices: ", s3)
}
