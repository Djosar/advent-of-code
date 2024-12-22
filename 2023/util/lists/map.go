package lists

// Map applies a function to each element of a slice and returns a new slice with the results.
func Map[T any, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}
