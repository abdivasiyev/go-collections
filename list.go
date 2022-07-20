package collections

type List[T Object] interface {
	Collection[T]
	Get(index int) (T, error)
	Set(index int, item T) error
	Delete(index int) error
}
