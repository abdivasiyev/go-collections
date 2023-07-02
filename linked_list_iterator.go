package collections

type LinkedListIterator[T Object] struct {
	list         List[T]
	currentIndex int
}

func NewLinkedListIterator[T Object](linkedList LinkedList[T]) *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		list:         &linkedList,
		currentIndex: 0,
	}
}

func (itr *LinkedListIterator[T]) HasNext() bool {
	return itr.currentIndex < itr.list.Size()
}

func (itr *LinkedListIterator[T]) Next() T {
	if itr.currentIndex >= itr.list.Size() || itr.currentIndex < 0 {
		panic(ErrIndexOutOfBounds)
	}

	item, err := itr.list.Get(itr.currentIndex)
	if err != nil {
		panic(err)
	}

	itr.currentIndex++

	return item
}
