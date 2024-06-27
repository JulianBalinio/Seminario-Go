package list

import (
	"Ej1/student"
	"errors"
)

var ErrEmpty = errors.New("empty list")

type Node struct {
	student *student.Student
	next    *Node
}

type List *Node

func New() List {
	return nil
}

func IsEmpty(self List) bool {
	return self == nil
}

func Len(self List) int {
	if IsEmpty(self) {
		return 0
	}
	return 1 + Len(Next(self))
}

func FrontElement(self List) *student.Student {
	return self.student
}

func Next(self List) List {
	return self.next
}

func ToString(self List) string {
	if IsEmpty(self) {
		return "{}"
	}
	return self.student.ToString() + "->" + ToString(Next(self))
}

func PushFront(self *List, std *student.Student) {
	aux := &Node{
		student: std,
		next:    *self,
	}
	*self = aux
}

func PushBack(self *List, std *student.Student) {
	if IsEmpty(*self) {
		*self = &Node{
			student: std,
			next:    nil,
		}
	} else {
      PushBack((*List) (&((*self).next)), std)
	}
}

func Remove(self *List) (*student.Student, error) {
	if IsEmpty(*self) {
		return nil, ErrEmpty
	}
	std := (*self).student
	*self = (*self).next
	return std, nil
}

func Iterate(self List, f func(*student.Student)) {
	aux := self
	for aux != nil {
		f(aux.student)
		aux = Next(aux)
	}
}
