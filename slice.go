// package slice is a collection of useful helpers for working more efficient
// with slices in go.
// The motivation behind this package has mostly been too experiment with generics
// released in go 1.18.
package slice

// Destruct1 returns first element in a slice
//	a := slice.Destruct1([]int{1, 2, 3, 4})
func Destruct1[A any](input []A) (a A) {
	if len(input) == 0 {
		return a
	}
	return input[0]
}

// Destruct2 returns the first two elements in a slice
//	a, b := slice.Destruct2([]int{1, 2, 3, 4})
func Destruct2[A any](input []A) (a, b A) {
	if len(input) == 0 {
		return a, b
	}
	if len(input) == 1 {
		return input[0], b
	}
	return input[0], input[1]
}

// Destruct3 returns the first three elements in a slice
//	a, b, c := slice.Destruct3([]int{1, 2, 3, 4})
func Destruct3[A any](input []A) (a, b, c A) {
	if len(input) == 0 {
		return a, b, c
	}
	if len(input) == 1 {
		return input[0], b, c
	}
	if len(input) == 2 {
		return input[0], input[1], c
	}
	return input[0], input[1], input[2]
}

// Destruct4 returns the first four elements in a slice
//	a, b, c, d := slice.Destruct4([]int{1, 2, 3, 4})
func Destruct4[A any](input []A) (a, b, c, d A) {
	if len(input) == 0 {
		return a, b, c, d
	}
	if len(input) == 1 {
		return input[0], b, c, d
	}
	if len(input) == 2 {
		return input[0], input[1], c, d
	}
	if len(input) == 3 {
		return input[0], input[1], input[2], d
	}
	return input[0], input[1], input[2], input[3]
}

// Merge two slices
func Merge[A any](a, b []A) []A {
	return append(a, b...)
}

// Find a value in a slice
func Find[A any](input []A, filter func(A) bool) (output A, index int) {
	for i := range input {
		if filter(input[i]) {
			return input[i], i
		}
	}
	return output, -1
}

// Contains a value in slice
func Contains[A any](input []A, filter func(A) bool) bool {
	for i := range input {
		if filter(input[i]) {
			return true
		}
	}
	return false
}

// Has value in slice
func Has[A any](input []A, filter func(A) bool) bool {
	return Contains(input, filter)
}

// Filter a slice of values
func Filter[A any](input []A, filter func(A) bool) (output []A) {
	for i := range input {
		if filter(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

// Map a slice of values
func Map[A, B any](input []A, mapper func(A) B) []B {
	output := make([]B, len(input))
	for i := range input {
		output[i] = mapper(input[i])
	}
	return output
}

// Reduce a slice of values
func Reduce[A, B any](input []A, output B, reducer func(A, B) B) B {
	for i := range input {
		output = reducer(input[i], output)
	}
	return output
}
