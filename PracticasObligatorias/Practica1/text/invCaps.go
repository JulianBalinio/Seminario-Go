package text

/*
	InvertCapitalization recorre la cadena de texto 's' y, si encuentra la palabra 'word',
	invierte la capitalizacion de cada char
	- Recibe dos parametros: 's' string original y 'word' la palabra a buscar y cambiar capitalizacion
	- Si la palabra 'word' se encuentra en la cadena, invierte la capitalización de cada letra en esa palabra
	- Retorna la cadena de texto con la capitalización invertida si la palabra 'word' se encuentra, de lo contrario,
	  retorna la cadena original
*/
func InvertCapitalization(s, word string) string {
    var modRunes []rune

    wordRunes := []rune(word)
    runes := []rune(s)

    wordLen := len(wordRunes)
    runesLen := len(runes)

    for i := 0; i < runesLen; {
        // Compruebo si es posible crear un slice de la longitud de la palabra a comparar
        if (i+wordLen <= runesLen) {

            wordSlice := runes[i : i+wordLen]

            // Compruebo si los slices son iguales
            if Equals(wordRunes, wordSlice) {
                modRunes = append(modRunes, invertCapitalization(wordSlice)...)
                i += len(wordSlice) // Avanzo al siguiente char despues de la palabra cambiada
            } else {
                modRunes = append(modRunes, runes[i]) // Si no hay coincidencia agrego el char actual
                i++
            }
        } else {
            // Si no es posible construir un slice de la longitud de la palabra a comparar, agrego el char
            modRunes = append(modRunes, runes[i])
            i++
        }
    }
    return string(modRunes)
}

/*
	Se recorre la cadena y se evalua la capitalizacion del char
	- Recibe un slice de runas de la palabra a cambiar capitalizacion
	- Se evalua la capitalizacion de cada runa y se invierte
	- Se retorna slice de runas con la capitalizacion invertida de la palabra
*/
func invertCapitalization(s []rune) []rune {
    var modRunes []rune
    for _, char := range s {
        if isUpper(char) {
            modRunes = append(modRunes, toLower(char))
        } else if isLower(char) {
            modRunes = append(modRunes, toUpper(char))
        } else {
            modRunes = append(modRunes, char) // Agrego el char no alfabetico
        }
    }
    return modRunes
}