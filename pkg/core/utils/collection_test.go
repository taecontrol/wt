package utils

import (
	"testing"
)

func TestNewCollection(t *testing.T) {
	t.Run("It creates a new collection", func(t *testing.T) {
		collection := NewCollection[int]([]int{1, 2, 3})

		if len(collection.Items) != 3 {
			t.Error("NewCollection should create a new collection")
		}

		if collection.Items[0] != 1 {
			t.Error("NewCollection should create a new collection")
		}

		if collection.Items[1] != 2 {
			t.Error("NewCollection should create a new collection")
		}

		if collection.Items[2] != 3 {
			t.Error("NewCollection should create a new collection")
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("It filters the items in the collection", func(t *testing.T) {
		collection := Collection[int]{Items: []int{1, 2, 3, 4, 5}}

		filteredCollection := collection.Filter(func(item int, index int) bool {
			return item%2 == 0
		})

		if len(filteredCollection.Items) != 2 {
			t.Error("Filter should filter the items in the collection")
		}

		if filteredCollection.Items[0] != 2 {
			t.Error("Filter should filter the items in the collection")
		}

		if filteredCollection.Items[1] != 4 {
			t.Error("Filter should filter the items in the collection")
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("It maps the items in the collection", func(t *testing.T) {
		collection := Collection[int]{Items: []int{1, 2, 3}}

		mappedCollection := collection.Map(func(item int, index int) int {
			return item * 2
		})

		if len(mappedCollection.Items) != 3 {
			t.Error("Map should map the items in the collection")
		}

		if mappedCollection.Items[0] != 2 {
			t.Error("Map should map the items in the collection")
		}

		if mappedCollection.Items[1] != 4 {
			t.Error("Map should map the items in the collection")
		}

		if mappedCollection.Items[2] != 6 {
			t.Error("Map should map the items in the collection")
		}
	})
}

func TestSlice(t *testing.T) {
	t.Run("It slices the items in the collection", func(t *testing.T) {
		collection := Collection[int]{Items: []int{1, 2, 3}}

		slicedCollection := collection.Slice(1, 3)

		if len(slicedCollection.Items) != 2 {
			t.Error("Slice should slice the items in the collection")
		}

		if slicedCollection.Items[0] != 2 {
			t.Error("Slice should slice the items in the collection")
		}

		if slicedCollection.Items[1] != 3 {
			t.Error("Slice should slice the items in the collection")
		}
	})
}

func TestCount(t *testing.T) {
	t.Run("It returns the number of items in the collection", func(t *testing.T) {
		collection := Collection[int]{Items: []int{1, 2, 3}}

		count := collection.Count()

		if count != 3 {
			t.Error("Count should return the number of items in the collection")
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("It reduces the items in the collection", func(t *testing.T) {
		collection := Collection[int]{Items: []int{1, 2, 3}}

		reduced := Reduce[int, int](collection, func(accumulator int, item int, index int) int {
			return accumulator + item
		}, 0)

		if reduced != 6 {
			t.Error("Reduce should reduce the items in the collection")
		}
	})
}
