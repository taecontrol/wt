package utils

type Collection[T any] struct {
	Items []T
}

type (
	FilterCallback[T any]        func(item T, index int) bool
	MapCallback[T any]           func(T, int) T
	ReduceCallback[T any, K any] func(K, T, int) K
)

func NewCollection[T any](items []T) Collection[T] {
	return Collection[T]{Items: items}
}

func (c *Collection[T]) Get(index int) T {
	return c.Items[index]
}

func (c *Collection[T]) Filter(callback FilterCallback[T]) Collection[T] {
	var filteredItems []T

	for index, item := range c.Items {
		if callback(item, index) {
			filteredItems = append(filteredItems, item)
		}
	}

	return Collection[T]{Items: filteredItems}
}

func (c *Collection[T]) Map(callback MapCallback[T]) Collection[T] {
	var mappedItems []T

	for index, item := range c.Items {
		mappedItem := callback(item, index)
		mappedItems = append(mappedItems, mappedItem)
	}

	return Collection[T]{
		Items: mappedItems,
	}
}

func (c *Collection[T]) Slice(start int, end int) Collection[T] {
	return Collection[T]{
		Items: c.Items[start:end],
	}
}

func (c *Collection[T]) Count() int {
	return len(c.Items)
}

func Reduce[T any, K any](c Collection[T], callback ReduceCallback[T, K], initialValue K) K {
	var accumulator = initialValue

	for index, item := range c.Items {
		accumulator = callback(accumulator, item, index)
	}

	return accumulator
}
