package main

import (
	"fmt"
	"sort"
)

func main() {
	// var arr = []string{}
	arr_2 := []int{5, 312, 32, 322}
	// fmt.Println(arr, reflect.TypeOf(arr), arr_2, len(arr_2))
	arr_2 = append(arr_2, 80980, 90)
	fmt.Print(arr_2, len(arr_2), "\n")

	arr_2 = append(arr_2[1:], 313, 9013) // this removes the 1st element as expected
	fmt.Print(arr_2, "\n")

	arr_2 = append(arr_2[1:4], 813, 113) // this removes the 1st element and all the elements from 4th index as expected
	fmt.Print(arr_2, "\n")

	// using the make and new keywords
	scores := make([]int, 5)
	fmt.Println(scores)

	scores[0] = 53
	scores[4] = 190
	scores[2] = 17
	sorted := sort.IntsAreSorted(scores)
	fmt.Println(sorted)
	fmt.Println(scores)
	// scores[5] = 89, this would give us an error

	scores = append(scores, 52, 32, 901, 3132) // doing this let's the array resize and accomodate all the values instead of giving an error
	fmt.Println(scores)

	sort.Ints(scores)
	fmt.Println(scores)
	sorted = sort.IntsAreSorted(scores)
	fmt.Println(sorted)

	// remove values from slice(index-based), we need to do this using append
	index := 3
	scores = append(scores[:index], scores[index+1:]...) // when appending 2 slices we need to spread the elements last slice, we can't concatenate more than two slices
	fmt.Println(scores)

	fmt.Println(scores[3:4])
}
