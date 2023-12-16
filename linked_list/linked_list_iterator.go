package linked_list

import "github.com/abdivasiyev/go-collections"

type linkedListIterator[T collections.Object] struct {
	list         collections.List[T]
	currentIndex int
}

func newLinkedListIterator[T collections.Object](linkedList LinkedList[T]) *linkedListIterator[T] {
	return &linkedListIterator[T]{
		list:         &linkedList,
		currentIndex: 0,
	}
}

func (itr *linkedListIterator[T]) HasNext() bool {
	return itr.currentIndex < itr.list.Size()
}

func (itr *linkedListIterator[T]) Next() T {
	if itr.currentIndex >= itr.list.Size() || itr.currentIndex < 0 {
		panic(collections.ErrIndexOutOfBounds)
	}

	item, err := itr.list.Get(itr.currentIndex)
	if err != nil {
		panic(err)
	}

	itr.currentIndex++

	return item
}
