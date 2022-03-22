package slice_test

import (
	"fmt"

	"github.com/kvartborg/slice"
)

func ExampleDesctruct1() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Destruct1(s),
	)
	// Output: 1
}

func ExampleDesctruct2() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Destruct2(s),
	)
	// Output: 1 2
}

func ExampleDesctruct3() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Destruct3(s),
	)
	// Output: 1 2 3
}

func ExampleDesctruct4() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Destruct4(s),
	)
	// Output: 1 2 3 4
}

func ExampleMerge() {
	a, b := []int{1, 2}, []int{3, 4}
	fmt.Println(
		slice.Merge(a, b),
	)
	// Output: [1 2 3 4]
}

func ExampleFind() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Find(s, func(value int) bool {
			return value == 3
		}),
	)
	// Output: 3 2
}

func ExampleContains() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Contains(s, func(value int) bool {
			return value == 4
		}),
	)
	// Output: true
}

func ExampleHas() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Has(s, func(value int) bool {
			return value == 4
		}),
	)
	// Output: true
}

func ExampleFilter() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Filter(s, func(value int) bool {
			return value%2 == 0
		}),
	)
	// Output: [2 4]
}

func ExampleMap() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(
		slice.Map(s, func(value int) string {
			return fmt.Sprint(value * 2)
		}),
	)
	// Output: [2 4 6 8 10]
}

func ExampleReduce() {
	s := []int{1, 2, 3, 4, 5}
	result := slice.Reduce(s, 0, func(value, sum int) int {
		return sum + value
	})
	fmt.Println(result)
	// Output: 15
}
