package list

type LinkedList interface {
	New() List
	IsEmpty() bool
	Len() int
	FrontElement(List) (int, error)
	Next() List
	ToString() (string, error)
	PushFront(int)
	PushBack(int)
	Remove() (int, error)
	Iterate(func(int) (int, error))
}