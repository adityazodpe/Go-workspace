package day3assignment

import (
	"fmt"
	"sort"
)

func Assign4() {
	fmt.Println("\nAssignment 4:")
	num := []int{84, 45, 61, 15, 72}
	fmt.Println("\n", num)

	sort.Ints(num)

	if sort.IntsAreSorted(num) == true {
		fmt.Println("\nInts are sorted.")
		fmt.Println("Sorted integers are: ", num)
	}

	val := []float64{2.5, 6.3, 1.5, 9.8, 4.7}
	fmt.Println("\n", val)

	sort.Float64s(val)

	if sort.Float64sAreSorted(val) == true {
		fmt.Println("\nFloats are sorted.")
		fmt.Println("Sorted floats are: ", val)
	}

	text := []string{"US", "India", "Germany", "Australia", "Spain"}
	fmt.Println("\n", text)

	sort.Strings(text)

	if sort.StringsAreSorted(text) == true {
		fmt.Println("\nStrings are sorted.")
		fmt.Println("Sorted strings are: ", text)
	}

}
