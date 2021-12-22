package main

import "fmt"


type numbers []int64

func newNumbers() numbers {
    intValues := []int64{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
	return intValues
}

func (n numbers) evenOrUneven() {
	for _, int := range n {
		if int %2 == 0 {
			fmt.Println(int, "is an even number")
		} else {
			fmt.Println(int, "is an odd number")
		}
	}
}
