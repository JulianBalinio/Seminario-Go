package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var upperCaseCount, lowerCaseCount, specialCharCount int
	numberCountMap := make(map[int]int)

	
	for i := 0; i <= 9; i++ {
		numberCountMap[i] = 0
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese la secuencia de caracteres (finalice con ENTER):")
	input, _ := reader.ReadString('\r')
	
	
	input = input[:len(input)-1]

	runes := []rune(input)
	filter(runes, &upperCaseCount, &lowerCaseCount, &specialCharCount, numberCountMap)

	fmt.Println("Cantidad de letras mayúsculas:", upperCaseCount)
	fmt.Println("Cantidad de letras minúsculas:", lowerCaseCount)
	fmt.Println("Cantidad de caracteres especiales:", specialCharCount)
	fmt.Println("Conteo de cada dígito decimal:", numberCountMap)
}

func filter(runes []rune, upperCaseCount, lowerCaseCount, specialCharCount *int, numberCountMap map[int]int) {
	for _, r := range runes {
		if unicode.IsUpper(r) {
			*upperCaseCount++
		} else if unicode.IsLower(r) {
			*lowerCaseCount++
		} else if unicode.IsDigit(r) {
			digit := int(r - '0')
			numberCountMap[digit]++
		} else {
			*specialCharCount++
		}
	}
}
