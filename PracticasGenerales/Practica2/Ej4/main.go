package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

const N = 5

func main() {
	var arrayX [N]int
	var arrayY [N]int
	var arrayZ [N]int

	fmt.Println("Ingrese los valores de las tres series:")
	FillArray(&arrayX, &arrayY, &arrayZ)

	min, max := MaxMin(&arrayY)
	result := (Sum(&arrayX) - Product(&arrayZ)*(float64(max)*float64(min)))

	fmt.Printf("El resultado del problema es --> %.2f\n", result)
}

var errRead = errors.New("error de lectura")

func read(valor *int) error {
	_, err := fmt.Scanln(valor)
	if err != nil {
		fmt.Println(errRead)
		return err
	}
	return nil
}

func FillArray(arrayX, arrayY, arrayZ *[N]int) {
	var valor int
	for i := 0; i < N; i++ {
		fmt.Printf("Valor %d para X: ", i+1)
		if read(&valor) != nil {
			return
		}
		arrayX[i] = valor
		fmt.Printf("Valor %d para Y: ", i+1)
		if read(&valor) != nil {
			return
		}
		arrayY[i] = valor
		fmt.Printf("Valor %d para Z: ", i+1)
		if read(&valor) != nil {
			return
		}
		arrayZ[i] = valor
	}
}

func Sum(arrayX *[N]int) float64 {
	var result float64
	for i := 0; i < N; i++ {
		if arrayX[i] != 0 {
			result += 1 / float64(arrayX[i])
		}
	}
	return result
}

func Product(arrayZ *[N]int) float64 {
	var result float64
	for i := 0; i < N; i++ {
		result += math.Pow(float64(arrayZ[i]), 3.0)
	}
	return result
}

func MaxMin(arrayY *[N]int) (int, int) {
	sliceY := arrayY[:]
	sort.Ints(sliceY)
	min := sliceY[0]
	max := sliceY[len(sliceY)-1]
	return min, max
}
