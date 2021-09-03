package main

import (
	"fmt"
)

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	size := 0
	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&size)
	arr := make([]int64, size)
	fmt.Println("Enter the inputs")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}

	max := arr[size-1]
	arr[size-1] = -1
	for i := size - 2; i >= 0; i-- {
		current := arr[i]
		arr[i] = max
		max = Max(max, current)
	}

	fmt.Println(arr)

}
