package text

// Resta entre valores ASCII
const (
    upperToLower = 'a' - 'A' // 32
    lowerToUpper = 'A' - 'a' // -32
)

// Control de capitalizacion

func isUpper(char rune) bool {
    return ((char >= 'A' && char <= 'Z') || (char >= 'Á' && char <= 'Ú') || (char == 'Ñ'))
}

func isLower(char rune) bool {
    return ((char >= 'a' && char <= 'z') || (char >= 'á' && char <= 'ú') || (char == 'ñ')) 
}

// Convertir a minuscula
func toLower(char rune) rune {
    if isUpper(char) {
        return char + upperToLower
    } 
    return char
}

// Convertir a mayuscula
func toUpper(char rune) rune {
    if isLower(char) {
        return char + lowerToUpper
    }
    return char
}

// Compruebo si las runas son iguales
func Equals(z []rune, y []rune) bool {
	for i := 0; i < len(z); i++ {
        if toLower(z[i]) != toLower(y[i]) {
            return false
        }
    }
	return true
}

