package array_list

import (
	"github.com/abdivasiyev/go-collections"
)

type (
	ArrayList[T collections.Object] struct {
		items []T
	}

	iterator[T collections.Object] interface {
		Reduce(reducer func(prev T, item T) T) T
		Skip(n int) iterator[T]
		SkipWhile(predicate func(item T) bool) iterator[T]
		Enumerate(enumerator func(index int, item T) T) iterator[T]
		Filter(predicate func(item T) bool) iterator[T]
		Map(mapper func(item T) T) iterator[T]
		Collect() *ArrayList[T]
	}

	intoIter[T collections.Object] struct {
		list *ArrayList[T]
	}
)

func (iter *intoIter[T]) Reduce(reducer func(prev T, item T) T) T {
	var (
		it       = iter.list.Iterator()
		prevItem T
	)

	for it.HasNext() {
		curr := it.Next()
		prevItem = reducer(prevItem, curr)
	}

	return prevItem
}

func (iter *intoIter[T]) Skip(n int) iterator[T] {
	var (
		list = New[T]()
		it   = iter.list.Iterator()
		i    = 0
	)

	for it.HasNext() {
		if i+1 < n {
			it.Next()
			continue
		}
		list.Add(it.Next())
		i++
	}

	iter.list = list
	return iter
}

func (iter *intoIter[T]) SkipWhile(predicate func(item T) bool) iterator[T] {
	var (
		list = New[T]()
		it   = iter.list.Iterator()
	)

	for it.HasNext() {
		item := it.Next()

		if predicate(item) {
			it.Next()
			continue
		}
		list.Add(it.Next())
	}

	iter.list = list
	return iter
}

func (iter *intoIter[T]) Enumerate(f func(index int, item T) T) iterator[T] {
	var (
		list = New[T]()
		it   = iter.list.Iterator()
		i    = 0
	)

	for it.HasNext() {
		enumerated := f(i, it.Next())
		list.Add(enumerated)
		i++
	}

	iter.list = list
	return iter
}

func (iter *intoIter[T]) Filter(f func(item T) bool) iterator[T] {
	list := New[T]()
	it := iter.list.Iterator()

	for it.HasNext() {
		val := it.Next()

		if f(val) {
			list.Add(val)
		}
	}

	iter.list = list
	return iter
}

func (iter *intoIter[T]) Map(f func(item T) T) iterator[T] {
	list := New[T]()
	it := iter.list.Iterator()

	for it.HasNext() {
		val := it.Next()
		list.Add(f(val))
	}

	iter.list = list
	return iter
}

func (iter *intoIter[T]) Collect() *ArrayList[T] {
	return iter.list
}

func New[T collections.Object]() *ArrayList[T] {
	return &ArrayList[T]{
		items: make([]T, 0),
	}
}

func (a *ArrayList[T]) ForEach() iterator[T] {
	return &intoIter[T]{list: a}
}

func (a *ArrayList[T]) WithSize(size int) *ArrayList[T] {
	a.items = make([]T, size)
	return a
}

func (a *ArrayList[T]) WithCapacity(capacity int) *ArrayList[T] {
	a.items = make([]T, 0, capacity)
	return a
}

func (a *ArrayList[T]) FromArray(items []T) {
	for _, item := range items {
		a.Add(item)
	}
}

func (a *ArrayList[T]) IsEmpty() bool {
	return len(a.items) == 0
}

func (a *ArrayList[T]) Size() int {
	return len(a.items)
}

func (a *ArrayList[T]) Iterator() collections.Iterator[T] {
	return newArrayListIterator[T](*a)
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

func (a *ArrayList[T]) AddAll(c collections.Iterable[T]) {
	itr := c.Iterator()

	for itr.HasNext() {
		a.Add(itr.Next())
	}
}

func (a *ArrayList[T]) RemoveAll(c collections.Iterable[T]) {
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
