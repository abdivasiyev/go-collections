package list

import "github.com/abdivasiyev/go-collections/collections"

type List[T collections.Object] interface {
	collections.Collection[T]
	Get(index int) (T, error)
	Set(index int, item T) error
	Delete(index int) error
}
