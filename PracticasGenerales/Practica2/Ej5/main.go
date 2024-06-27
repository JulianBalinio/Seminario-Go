package main

import (
	"errors"
	"fmt"
)

const ArraySize = 10

type Array [ArraySize]float64

func main() {
	var arr1, arr2 Array
	var f float64
	
	fmt.Println("Introduzca valores para inicializar el array:")
	Initialize(&arr1, f)
	Initialize(&arr2, f)
	
	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	
	sum := Sum(arr1, arr2)
	fmt.Println("Suma de arrays:", sum)
	
	SumInPlace(&arr1, arr2)
	fmt.Println("Array 1 post SumInPlace:", arr1)
}

var errRead = errors.New("reading error")

func read(value *float64) error {
	_, err := fmt.Scanln(value)
	if err != nil {
		fmt.Println(errRead)
		return err
	}
	return nil
}

func Initialize(arr *Array, f float64) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Enter value %d: ", i+1)
		if read(&f) != nil {
			return
		}
		arr[i] = f
	}
}

func Sum(arr1, arr2 Array) Array {
	var summed Array
	for i := 0; i < len(arr1); i++ {
		summed[i] = arr1[i] + arr2[i]
	}
	return summed
}

func SumInPlace(arr1 *Array, arr2 Array) {
	for i := 0; i < len(arr1); i++ {
		arr1[i] += arr2[i]
	}
}
