package main

import (
	"Practica1/text"
    "fmt"
)

/*
    - Tanto ReplaceWord(conv.go) e InvertCapitalization(invert.go) mantienen la misma logica
    - go.mod para poder usar dependencias locales
*/

func main() {
    
    // Ejercicio 1 y 2
    //fmt.Print("Ingrese una frase/texto: ")
    //text.ReadText(&sentence) 
    sentence := "el MiÉRcoLes salgo con mi AutomÓViL"
    fWord := "miércoles"
    sWord := "automóvil"

    result := text.ReplaceWord(sentence, fWord, sWord)
    fmt.Println(result)

    // Ejercicio 3
    sentence = "pequeño PEQUEÑO peQUEño PEqueÑo."
    palabra := "PeQUeÑo"
    resultado := text.InvertCapitalization(sentence, palabra)
    fmt.Println(resultado)
}