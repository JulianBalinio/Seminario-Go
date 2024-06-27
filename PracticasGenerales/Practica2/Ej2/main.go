package main

import "fmt"

func main(){

	for i:= 0; i < 20; i++ {
		resultRecursive := recursiveFactorial(i)
		resultIterative := iterativeFactorial(i)
		fmt.Println(resultRecursive)
		fmt.Println(resultIterative)
	}
}

func recursiveFactorial(n int) int {
	var result int
	if (n == 0) {
		return 1
	} else if (n > 0){
		result = n * recursiveFactorial(n-1)
	}
	return result
}

func iterativeFactorial(n int) int {
	var result int
	if (n == 0) {
		return 1
	} else {
		result = 1
		for i:= 1; i < n; i++ {
			result *= (i+1)
		}
	}
	return result
}