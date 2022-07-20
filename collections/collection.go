package collections

import (
	"errors"
)

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds")
)

type Object interface {
	comparable
}

type Collection[T Object] interface {
	IsEmpty() bool
	Size() int
	Iterator() Iterator[T]
	ToArray() []T
	Add(item T)
	Remove(item T)
	Contains(item T) bool
	Clear()
	AddAll(c Collection[T])
	RemoveAll(c Collection[T])
}
