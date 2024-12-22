package strings

// Reverse returns the reversed string of the input.
func Reverse(val string) (result string) {
	for _, v := range val {
		result = string(v) + result
	}
	return result
}
