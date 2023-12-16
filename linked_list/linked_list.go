package linked_list

import "github.com/abdivasiyev/go-collections"

type nodeValue[T collections.Object] struct {
	value T
	next  *nodeValue[T]
}

type LinkedList[T collections.Object] struct {
	head *nodeValue[T]
	tail *nodeValue[T]
	size int
}

func New[T collections.Object]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (a *LinkedList[T]) FromArray(items []T) {
	for _, item := range items {
		a.Add(item)
	}
}

func (a *LinkedList[T]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *LinkedList[T]) Size() int {
	return a.size
}

func (a *LinkedList[T]) Iterator() collections.Iterator[T] {
	return newLinkedListIterator[T](*a)
}

func (a *LinkedList[T]) ToArray() []T {
	var (
		result = make([]T, 0)
		itr    = a.Iterator()
	)

	for itr.HasNext() {
		result = append(result, itr.Next())
	}

	return result
}

func (a *LinkedList[T]) Add(item T) {
	a.size++
	if a.size == 1 {
		a.head = &nodeValue[T]{
			value: item,
		}
		a.tail = a.head
		return
	}
	tailNode := a.tail
	a.tail = &nodeValue[T]{value: item}
	tailNode.next = a.tail
}

func (a *LinkedList[T]) Remove(item T) {
	foundIndex := a.findItem(item)
	if foundIndex == -1 {
		return
	}

	a.Delete(foundIndex)
}

func (a *LinkedList[T]) Contains(item T) bool {
	return a.findItem(item) != -1
}

func (a *LinkedList[T]) Clear() {
	a.size = 0
	a.head = nil
	a.tail = nil
}

func (a *LinkedList[T]) AddAll(c collections.Collection[T]) {
	itr := c.Iterator()

	for itr.HasNext() {
		a.Add(itr.Next())
	}
}

func (a *LinkedList[T]) RemoveAll(c collections.Collection[T]) {
	itr := c.Iterator()

	for itr.HasNext() {
		a.Remove(itr.Next())
	}
}

func (a *LinkedList[T]) Get(index int) (T, error) {
	var result T
	if index < 0 || index >= a.size {
		return result, collections.ErrIndexOutOfBounds
	}

	var node *nodeValue[T] = a.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value, nil
}

func (a *LinkedList[T]) Set(index int, item T) error {
	if index < 0 || index >= a.size {
		return collections.ErrIndexOutOfBounds
	}

	var node *nodeValue[T] = a.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	node.value = item

	return nil
}

func (a *LinkedList[T]) Delete(index int) error {
	if index < 0 || index >= a.size {
		return collections.ErrIndexOutOfBounds
	}

	var node *nodeValue[T] = a.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	node.next = node.next.next
	a.size--

	return nil
}

func (a *LinkedList[T]) findItem(item T) int {
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
