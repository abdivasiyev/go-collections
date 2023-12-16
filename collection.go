package collections

import (
	"errors"
)

var (
	ErrIndexOutOfBounds = errors.New("index out of bounds") // returns when index not satisfied by collection
)

// Object comparable interface for working with any comparable data type in go
type Object interface {
	comparable
}

// Collection base interface for all data structures
type Collection[T Object] interface {
	IsEmpty() bool             // returns true if collection is empty
	Size() int                 // returns current size of collection
	ToArray() []T              // converts collection to slice
	FromArray(items []T)       // converts slice of items to specific collection
	Add(item T)                // adds new item to collection
	Remove(item T)             // removes given item from collection
	Contains(item T) bool      // checks given item exists in collection
	Clear()                    // clears collection
	AddAll(c Collection[T])    // adds given collection items to current collection
	RemoveAll(c Collection[T]) // removes given collection items from current collection
}
