package array_list

import "github.com/abdivasiyev/go-collections"

type arrayListIterator[T collections.Object] struct {
	list         collections.List[T]
	currentIndex int
}

func newArrayListIterator[T collections.Object](arrayList ArrayList[T]) *arrayListIterator[T] {
	return &arrayListIterator[T]{
		list:         &arrayList,
		currentIndex: 0,
	}
}

func (itr *arrayListIterator[T]) HasNext() bool {
	return itr.currentIndex < itr.list.Size()
}

func (itr *arrayListIterator[T]) Next() T {
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
