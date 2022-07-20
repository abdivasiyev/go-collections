package list

import (
	"github.com/abdivasiyev/go-collections/collections"
)

type ArrayList[T collections.Object] struct {
	items []T
}

func NewArrayList[T collections.Object]() *ArrayList[T] {
	return &ArrayList[T]{
		items: make([]T, 0),
	}
}

func (a *ArrayList[T]) IsEmpty() bool {
	return len(a.items) == 0
}

func (a *ArrayList[T]) Size() int {
	return len(a.items)
}

func (a *ArrayList[T]) Iterator() collections.Iterator[T] {
	return NewArrayListIterator[T](*a)
}

func (a *ArrayList[T]) ToArray() []T {
	var (
		result = make([]T, 0)
		itr    = a.Iterator()
	)

	for itr.HasNext() {
		result = append(result, itr.Next())
	}

	return result
}

func (a *ArrayList[T]) Add(item T) {
	a.items = append(a.items, item)
}

func (a *ArrayList[T]) Remove(item T) {
	foundIndex := a.findItem(item)
	if foundIndex == -1 {
		return
	}

	a.items = append(a.items[:foundIndex], a.items[foundIndex+1:]...)
}

func (a *ArrayList[T]) Contains(item T) bool {
	return a.findItem(item) != -1
}

func (a *ArrayList[T]) Clear() {
	a.items = nil
}

func (a *ArrayList[T]) AddAll(c collections.Collection[T]) {
	itr := c.Iterator()

	for itr.HasNext() {
		a.Add(itr.Next())
	}
}

func (a *ArrayList[T]) RemoveAll(c collections.Collection[T]) {
	itr := c.Iterator()

	for itr.HasNext() {
		a.Remove(itr.Next())
	}
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	var result T
	if index < 0 || index >= len(a.items) {
		return result, collections.ErrIndexOutOfBounds
	}

	return a.items[index], nil
}

func (a *ArrayList[T]) Set(index int, item T) error {
	if index < 0 || index >= len(a.items) {
		return collections.ErrIndexOutOfBounds
	}

	a.items[index] = item

	return nil
}

func (a *ArrayList[T]) Delete(index int) error {
	if index < 0 || index >= len(a.items) {
		return collections.ErrIndexOutOfBounds
	}

	a.items = append(a.items[:index], a.items[index+1:]...)
	return nil
}

func (a *ArrayList[T]) findItem(item T) int {
	var (
		index = 0
		itr   = a.Iterator()
	)

	for itr.HasNext() {
		if itr.Next() == item {
			return index
		}

		index++
	}

	return -1
}
