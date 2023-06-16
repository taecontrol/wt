package utils

type Mock struct {
	Count int
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) Call() {
	m.Count++
}

func (m *Mock) WasCalledTimes(count int) bool {
	return m.Count == count
}
