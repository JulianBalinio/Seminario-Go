package main

import (
	"fmt"
)

func main() {
	firstSlice := []int{81, 62, 11, 24, 15, 88, 101}
	secondSlice := []int{6, 26, 8, 21, 33, 111, 19, 13, 14, 109}
	
	fmt.Println("Slice 1:", firstSlice)
	fmt.Println("Slice 2:", secondSlice)
	
	summed := Sum(firstSlice, secondSlice)
	fmt.Println("Summed slice:", summed)
	
	average := Avg(secondSlice)
	fmt.Printf("Average of slice 2: %.2f\n", average)
}


func Sum(a, b []int) []int {
	var minLen int
	if len(a) < len(b) {
		minLen = len(a)
	} else {
		minLen = len(b)
	}
	sliceSum := make([]int, minLen)
	for i := 0; i < minLen; i++ {
		sliceSum[i] = a[i] + b[i]
	}
	return sliceSum
}

// Avg calculates and returns the average of the elements in the input slice.
func Avg(a []int) float64 {
	var sum int
	for _, value := range a {
		sum += value
	}
	return float64(sum) / float64(len(a))
}
