package genericsort

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortStabilityTestData = []emp{
	{
		fname: "Michael",
		lname: "Scott",
		age:   45,
		order: 1,
	},
	{
		fname: "Michael",
		lname: "Scott",
		age:   45,
		order: 9,
	},
	{
		fname: "Dwight",
		lname: "Schrute",
		age:   38,
		order: 2,
	},
	{
		fname: "Jim",
		lname: "Halpert",
		age:   35,
		order: 3,
	},
}

func Test_HeapSortOnIntegerSliceAscending(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out []int
	}{
		"non empty": {
			in:  []int{4, 99, 1, 2, 7, 10},
			out: []int{1, 2, 4, 7, 10, 99},
		},
		"empty": {
			in:  []int{},
			out: []int{},
		},
		"single entry": {
			in:  []int{4},
			out: []int{4},
		},
		"double entry": {
			in:  []int{4, 1},
			out: []int{1, 4},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.in
			k := HeapSort(a, func(i, j int) bool {
				return a[i] < a[j]
			})
			x := reflect.DeepEqual(test.out, k)
			assert.Equal(t, x, true)
		})
	}
}

func Test_HeapSortOnIntegerSliceDescending(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out []int
	}{
		"non empty": {
			in:  []int{4, 99, 1, 2, 7, 10},
			out: []int{99, 10, 7, 4, 2, 1},
		},
		"empty": {
			in:  []int{},
			out: []int{},
		},
		"single entry": {
			in:  []int{4},
			out: []int{4},
		},
		"double entry": {
			in:  []int{1, 4},
			out: []int{4, 1},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.in
			k := HeapSort(a, func(i, j int) bool {
				return a[i] > a[j]
			})
			x := reflect.DeepEqual(test.out, k)
			assert.Equal(t, x, true)
		})
	}
}

func Test_MergeSortOnIntegerSliceAscending(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out []int
	}{
		"non empty": {
			in:  []int{4, 99, 1, 2, 7, 10},
			out: []int{1, 2, 4, 7, 10, 99},
		},
		"empty": {
			in:  []int{},
			out: []int{},
		},
		"single entry": {
			in:  []int{4},
			out: []int{4},
		},
		"double entry": {
			in:  []int{4, 1},
			out: []int{1, 4},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.in
			k := MergeSort(a, func(i, j int) bool {
				return a[i] < a[j]
			})
			x := reflect.DeepEqual(test.out, k)
			assert.Equal(t, x, true)
		})
	}
}

func Test_MergeSortOnIntegerSliceDescending(t *testing.T) {
	tests := map[string]struct {
		in  []int
		out []int
	}{
		"non empty": {
			in:  []int{4, 99, 1, 2, 7, 10},
			out: []int{99, 10, 7, 4, 2, 1},
		},
		"empty": {
			in:  []int{},
			out: []int{},
		},
		"single entry": {
			in:  []int{4},
			out: []int{4},
		},
		"double entry": {
			in:  []int{1, 4},
			out: []int{4, 1},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.in
			k := MergeSort(a, func(i, j int) bool {
				return a[i] > a[j]
			})
			x := reflect.DeepEqual(test.out, k)
			assert.Equal(t, x, true)
		})
	}
}

func Test_MergeSortStabilityAscending(t *testing.T) {
	tests := map[string]struct {
		in  []emp
		out []emp
	}{
		"non empty": {
			in: sortStabilityTestData,
			out: []emp{
				{
					fname: "Dwight",
					lname: "Schrute",
					age:   38,
					order: 2,
				},
				{
					fname: "Jim",
					lname: "Halpert",
					age:   35,
					order: 3,
				},
				{
					fname: "Michael",
					lname: "Scott",
					age:   45,
					order: 1,
				},
				{
					fname: "Michael",
					lname: "Scott",
					age:   45,
					order: 9,
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.in
			k := MergeSort(a, func(i, j int) bool {
				return a[i].fname < a[j].fname
			})
			x := reflect.DeepEqual(test.out, k)
			assert.Equal(t, x, true)
		})
	}
}

func Test_MergeSortStabilityDescending(t *testing.T) {
	tests := map[string]struct {
		in  []emp
		out []emp
	}{
		"non empty": {
			in: sortStabilityTestData,
			out: []emp{
				{
					fname: "Michael",
					lname: "Scott",
					age:   45,
					order: 1,
				},
				{
					fname: "Michael",
					lname: "Scott",
					age:   45,
					order: 9,
				},
				{
					fname: "Jim",
					lname: "Halpert",
					age:   35,
					order: 3,
				},
				{
					fname: "Dwight",
					lname: "Schrute",
					age:   38,
					order: 2,
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := test.in
			k := MergeSort(a, func(i, j int) bool {
				return a[i].fname > a[j].fname
			})
			x := reflect.DeepEqual(test.out, k)
			assert.Equal(t, x, true)
		})
	}
}
