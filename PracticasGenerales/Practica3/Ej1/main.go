package main

import (
    "fmt"
)

//K es el tipo de la clave y debe ser comparable (es decir, soportar operaciones de comparación como ==)
//V es el tipo del valor y puede ser de cualquier tipo (any)
type Map[K comparable, V any] map[K]V


func main() {
    
	var ageMap Map[string, int] = Map[string, int]{
		"Julian": 27,
		"Agustina": 26,
	}

	fmt.Println("Age Map:")
	for key, value := range ageMap {
    	fmt.Printf("%s: %d\n", key, value)
	}

	////////////////////////////////////////////////////////

	var codeMap Map[int, string] = Map[int, string]{
		404: "Not Found",
		500: "Internal Server Error",
	}

	fmt.Println("Code Map:")
	for key, value := range codeMap {
    	fmt.Printf("%d: %s\n", key, value)
	}
}