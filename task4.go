package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func ReplaceMaxFromRight(arr []int) {
	size := len(arr)
	max := arr[size-1] // max = last element of array for a situation with negative elements
	arr[size-1] = -1   // last element = -1
	for i := size - 2; i >= 0; i-- {
		current := arr[i] //go through the array from right to left replacing the element with the maximum on the right
		arr[i] = max
		max = Max(max, current)
	}
}

func main() {
	size := 0
	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&size)
	arr := make([]int, size)
	fmt.Println("Enter the inputs")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}

	ReplaceMaxFromRight(arr)

	fmt.Println(arr)

}
