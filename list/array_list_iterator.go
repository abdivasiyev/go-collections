package list

import "github.com/abdivasiyev/go-collections/collections"

type ArrayListIterator[T collections.Object] struct {
	list         List[T]
	currentIndex int
}

func NewArrayListIterator[T collections.Object](arrayList ArrayList[T]) *ArrayListIterator[T] {
	return &ArrayListIterator[T]{
		list:         &arrayList,
		currentIndex: 0,
	}
}

func (itr *ArrayListIterator[T]) HasNext() bool {
	return itr.currentIndex < itr.list.Size()
}

func (itr *ArrayListIterator[T]) Next() T {
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
