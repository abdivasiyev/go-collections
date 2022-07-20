package collections

// Iterator loops through collection and returns items
type Iterator[T Object] interface {
	HasNext() bool // checks iterator has next item
	Next() T       // returns next element if exists or panics with ErrIndexOutOfBounds
}
