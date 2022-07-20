package collections

type Iterator[T Object] interface {
	HasNext() bool
	Next() T
}

type ListIterator[T Object] struct {
	list         List[T]
	currentIndex int
}

func NewListIterator[T Object](list List[T]) *ListIterator[T] {
	return &ListIterator[T]{
		list:         list,
		currentIndex: 0,
	}
}

func (l ListIterator[T]) HasNext() bool {
	//TODO implement me
	panic("implement me")
}

func (l ListIterator[T]) Next() T {
	//TODO implement me
	panic("implement me")
}
