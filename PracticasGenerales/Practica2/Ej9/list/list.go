package list

import (
	"strconv"
	"errors"
)

var EmptyError = errors.New("Empty list")
var EmptyNext = errors.New("Doesn't exist next element")
var InvalidTypeError = errors.New("Only integers are allowed")

type Node struct {
	data int
	next *Node
}

type List struct {
	elem *Node
	len int
}


//Instancia una lista y devuelvo un puntero
func New() *List {
	return &List{}
}

func (l *List) IsEmpty() bool {
	return l.elem == nil
}

func (l *List) Len() int {
	return l.len
}

func (l *List) FrontElement() (int, error) {
	if l.IsEmpty() {
		return 0, EmptyError
	}
	return l.elem.data, nil
}

func (l *List) Next() (int, error) {
	if (l.IsEmpty()) {
		return 0, EmptyError
	} else if(l.elem.next == nil) {
		return 0, EmptyNext
	}
	return l.elem.next.data, nil
}

func (l* List) ToString() (string, error){
	if l.IsEmpty() {
		return "", EmptyError
	}
	currentNode := l.elem
	myStr := "["
	for currentNode != nil {
		myStr += strconv.Itoa(currentNode.data) //Convertir int a string
		if(currentNode.next != nil) {
			myStr += ", "
		}	
		currentNode = currentNode.next
	}
	myStr += "]"
	return myStr, nil
}

func(l *List) PushFront(elem int) {
	newNode := &Node{data: elem, next: l.elem}
	l.elem = newNode
	l.len++
}

func (l* List) PushBack(elem int) {
	newNode := &Node{data: elem}
	if (l.IsEmpty()) {
		l.elem = newNode
	} else {
		currentNode := l.elem
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	l.len++
}


//Recorre los elementos de la lista y aplica una funcion anonima a cada dato del nodo
func (l *List) Iterate(f func(int) int) error {
	if(l.IsEmpty()) {
		return EmptyError
	}
    currentNode := l.elem
    for currentNode != nil {
        currentNode.data = f(currentNode.data)
        currentNode = currentNode.next
    }
	return nil
}

func (l *List) Remove() (int, error){
	if (l.IsEmpty()) {
		return 0, EmptyError
	}
	deletedData := l.elem.data
	l.elem = l.elem.next
	l.len--
	return deletedData, nil
}