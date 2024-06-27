package text

/*
	ReplaceWord recorre la cadena de texto 's' utilizando runas y reemplaza las ocurrencias de una palabra por otra
	- Recibe tres parametros: 's' string original, 'fWord' y 'sWord' son las palabras a intercambiar
	- Se recorre la string mediante runas para mejor manipulacion de char
	- Reemplaza las ocurrencias de 'fWord' por 'sWord' manteniendo la capitalizacion original
	- Retorna el arreglo de runas de la cadena original y modificado como tipo string
    Faltaria ajustar el codigo en caso de que las palabras tengas distinta len
        - Ya sea manejando un error 
*/
func ReplaceWord(s, fWord, sWord string) string {
    var modRunes []rune

    fWordRunes := []rune(fWord)
    sWordRunes := []rune(sWord)
    runes := []rune(s)

	fWordLen := len(fWordRunes)
	sWordLen := len(sWordRunes)
	runesLen := len(runes)
	
    for i := 0; i < runesLen; {
        // Comprobar si es posible crear un slice de la longitud de la palabra a comparar
        if (i+fWordLen <= runesLen) && (i+sWordLen <= runesLen) {

            wordFirstSlice := runes[i : i+fWordLen]
            wordSecondSlice := runes[i : i+sWordLen]

            // Comprobar si las runas coinciden
            if Equals(fWordRunes, wordFirstSlice) {
                modRunes = append(modRunes, replaceWithCapitalization(wordFirstSlice, sWordRunes)...) // ... para descomponer el slice retornado y agregarlo a modRunes
                i += len(wordFirstSlice) // Avanzo al siguiente char despues de la palabra reemplazada
            } else if Equals(sWordRunes, wordSecondSlice) {
                modRunes = append(modRunes, replaceWithCapitalization(wordSecondSlice, fWordRunes)...) 
                i += len(wordSecondSlice) // Avanzo al siguiente char despues de la palabra reemplazada
            } else {
                modRunes = append(modRunes, runes[i]) // Si no hay coincidencia agrego el char actual
                i++
            }
        } else {
            // Si no es posible constriur un slice de la long. de la palabra a comparar, agrego el char
            modRunes = append(modRunes, runes[i])
            i++
        }
    }
    return string(modRunes)
}

/*
	Se recorre la cadena y se evalua la capitalizacion del char
	- Recibe dos slices de runas de las palabras a cambiar
	- Se evalua la capitalizacion de cada runa
	- Se retorna slice de runas capitalizadas de la palabra buscada
*/
func replaceWithCapitalization(s []rune, wanted []rune) []rune {
    var modRunes []rune
    for i, char := range s {
        if i < len(wanted) {
            if isUpper(char) {
                modRunes = append(modRunes, toUpper(wanted[i]))
            } else {
                modRunes = append(modRunes, toLower(wanted[i]))
            }
        }
    }
    return modRunes
}