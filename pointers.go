package go_utils

// Ptr will convert passed value to pointer to passed value. Useful for getting pointer to literal values
func Ptr[A any](v A) *A {
	return &v
}
