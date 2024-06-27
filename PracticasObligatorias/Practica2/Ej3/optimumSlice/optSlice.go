package optimumSlice

import "errors"

type Pair struct {
	elem int
	quantity int
}

type OptimumSlice []Pair

var ErrEmptyOptimum = errors.New("OptimumSlice is empty")


func New(s []int) OptimumSlice {
	if len(s) == 0 {
		return OptimumSlice{}
	}

	optimumSlice := make(OptimumSlice, 0)
	current := s[0] //tomo el primer elemento del slice
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == current {
			count++ //se suma Pair.quantity si el num actual coincide con current
		} else {
			optimumSlice = append(optimumSlice, Pair{elem: current, quantity: count})
			current = s[i]
			count = 1
		}
	}
	optimumSlice = append(optimumSlice, Pair{elem: current, quantity: count})
	return optimumSlice
}

func IsEmpty(o OptimumSlice) bool {
	return Len(o) == 0
}

func Len(o OptimumSlice) int {
	length := 0
	for _, pair := range o {
		length += pair.quantity
	}
	return length
}

func FrontElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, ErrEmptyOptimum
	}
	return o[0].elem, nil
}

func LastElement(o OptimumSlice) (int, error) {
	if IsEmpty(o) {
		return 0, ErrEmptyOptimum
	}
	return o[Len(o)-1].elem, nil
}

func Insert(o OptimumSlice, element int, position int) OptimumSlice {
    if position < 0 || position > Len(o) {
        return o
    }

    newSlice := make(OptimumSlice, 0, len(o)+1)
    currentPosition := 0

    for i, pair := range o {
        nextPosition := currentPosition + pair.quantity
        if position <= nextPosition {
            //agrego el elemento dentro del par actual
            if element == pair.elem {
                pair.quantity++
                newSlice = append(newSlice, pair)
            } else {
                if position > currentPosition {
                    newSlice = append(newSlice, Pair{elem: pair.elem, quantity: position - currentPosition})
                }
                newSlice = append(newSlice, Pair{elem: element, quantity: 1})
                if position < nextPosition {
                    newSlice = append(newSlice, Pair{elem: pair.elem, quantity: nextPosition - position})
                }
            }
            newSlice = append(newSlice, o[i+1:]...)
            return newSlice
        }
        newSlice = append(newSlice, pair)
        currentPosition = nextPosition
    }

    //agrego al final si la posicion es igual a la longitud del slice
    newSlice = append(newSlice, Pair{elem: element, quantity: 1})
    return newSlice
}

func SliceArray(o OptimumSlice) []int {
	slice := make([]int, 0)
	for _, pair := range o {
		for i := 0; i < pair.quantity; i++ {
			slice = append(slice, pair.elem)
		}
	}
	return slice
}