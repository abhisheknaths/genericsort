package genericsort_test

import (
	"fmt"
	"testing"

	"github.com/abhisheknaths/genericsort"
)

/*
	This file is intended to provide an area for developers to play around with the package.
	Run go test -v -run Test_Main.
*/

func Test_Main(t *testing.T) {
	a := []int{9, 2, 1, 6, 3, 0, 4}
	// we can specify type param
	a = genericsort.HeapSort[int](a, func(i, j int) bool {
		return a[i] < a[j]
	})
	fmt.Println(a)

	b := []string{"f", "e", "r", "n", "a", "n", "d", "o"}
	// if we dont it is inferred
	b = genericsort.HeapSort(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	fmt.Println(b)
}
