package geo

type Matrix[T any] struct {
	Width, Height int
	Elements      []T
}

// NewMatrix creates a new matrix with the given width and height.
func NewMatrix[T any](width, height int) *Matrix[T] {
	return &Matrix[T]{
		Width:    width,
		Height:   height,
		Elements: make([]T, width*height),
	}
}

// Get returns the element at the given row and column.
func (m *Matrix[T]) Get(row, col int) T {
	return m.Elements[row*m.Width+col]
}

// Set sets the element at the given row and column.
func (m *Matrix[T]) Set(row, col int, value T) {
	m.Elements[row*m.Width+col] = value
}

// Fill sets all elements in the matrix to the given value.
func (m *Matrix[T]) Fill(value T) {
	for i := range m.Elements {
		m.Elements[i] = value
	}
}

// Transpose returns a new matrix that is the transpose of the current matrix.
func (m *Matrix[T]) Transpose() *Matrix[T] {
	transposed := NewMatrix[T](m.Height, m.Width)
	for row := 0; row < m.Height; row++ {
		for col := 0; col < m.Width; col++ {
			transposed.Set(col, row, m.Get(row, col))
		}
	}
	return transposed
}
