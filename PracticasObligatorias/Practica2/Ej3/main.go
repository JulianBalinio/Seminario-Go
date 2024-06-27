package main

import (
	"Ej3/optimumSlice"
	"fmt"
)

func main() {
	s := []int{3, 3, 3, 3, 3, 1, 1, 1, 1, 1, 1, 1, 1, 23, 23, 23, 23, 23, 23, 23, 23, 23, 3, 3, 3, 3, 3, 3, 3, 3, 7, 5, 5, 5}
	o := optimumSlice.New(s)
	fmt.Println("Original OptimumSlice:", o, " Len: ", optimumSlice.Len(o))

	element := 9
	position := 6
	o = optimumSlice.Insert(o, element, position)
	fmt.Println("OptimumSlice after insertion:", o, " Len: ", optimumSlice.Len(o))

	o = optimumSlice.Insert(o, 5, 11)

	fmt.Println("OptimumSlice after insertion:", o, " Len: ", optimumSlice.Len(o))

	fmt.Println("Converted slice:", optimumSlice.SliceArray(o))
}