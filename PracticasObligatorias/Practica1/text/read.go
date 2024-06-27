package text

import "fmt"

/*	Uso de runas para lectura estandar, esto evita problemas con char especiales, a diferencia del tipo byte
	- Esto se ve reflejado en el Ejercicio 2, cuando se incluyen tildes
	Recibe un puntero a una cadena donde se almacena la entrada
	El bucle corta cuando se ingresa un salto de linea(enter) o se cierra la consola(EoF)
*/
func ReadText(s *string) {

	var text string

	for {
		var char rune
		
		_, err:= fmt.Scanf("%c", &char)
		if err != nil {
			fmt.Println("Error inesperado: ", err)
			return
		}

		
		if (char == '\n'){
			break
		}
	
		text += string(char)
	}

    *s = text
}


