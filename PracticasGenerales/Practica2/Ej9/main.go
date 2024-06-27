package main

import (
	"fmt"
	"Ej9/list"
)

func main(){

	myList := list.New()

	fmt.Println(rune('a'))
	myList.PushFront('a')
	myList.PushFront(211)
	myList.PushBack(500)

	fmt.Println(myList.Len())

	if elem, err := myList.FrontElement(); err == nil {
		fmt.Println("Primer elemento de la lista:", elem)
	} else {
		fmt.Println("Error al buscar el primer elemento:", err)
	}

	if str, err := myList.ToString(); err == nil {
		fmt.Println("Lista como cadena:", str)
	} else {
		fmt.Println("Error al convertir en string:", err)
	}

	err := myList.Iterate(func(x int) int {
		return x*2
	})
	if err != nil {
		fmt.Println("Error al iterar:", err)
	}

	if str, err := myList.ToString(); err == nil {
		fmt.Println("Lista como cadena:", str)
	} else {
		fmt.Println("Error al convertir en string:", err)
	}
}