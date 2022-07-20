package collections

// List interface for listable data types
type List[T Object] interface {
	Collection[T]
	Get(index int) T
}

type ArrayList[T Object] struct {
	items []T
}

func NewArrayList[T Object]() *ArrayList[T] {
	return &ArrayList[T]{
		items: make([]T, 0),
	}
}

func (a *ArrayList[T]) Add(element T) bool {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) AddAll(elements []T) bool {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Remove(element T) bool {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) RemoveAll(elements []T) bool {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Contains(element T) bool {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) ToArray() []T {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Iterator() Iterator[T] {
	//TODO implement me
	panic("implement me")
}

func (a *ArrayList[T]) Get(index int) T {
	//TODO implement me
	panic("implement me")
}
