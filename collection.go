package collections

type Object interface {
	comparable
}

type Collection[T Object] interface {
	Add(element T) bool          // adds given element to collection
	AddAll(elements []T) bool    // adds given elements to collection
	Remove(element T) bool       // removes given element from collection
	RemoveAll(elements []T) bool // removes given elements from collection
	Size() int                   // returns collection current size
	Clear()                      // clears collection
	Contains(element T) bool     // checks collection has element
	ToArray() []T                // returns array of items in collection
	Iterator() Iterator[T]       // returns iterator for given collection
}
