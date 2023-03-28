package genericsort

import (
	"sort"
	"testing"
)

type emp struct {
	fname string
	lname string
	age   int
	order int
}

type sortMethod func(i, j int) bool

func getTestData() []emp {
	return []emp{
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
		{
			fname: "Kelly",
			lname: "Kapoor",
			age:   28,
			order: 2,
		},
		{
			fname: "Ryan",
			lname: "Howard",
			age:   30,
			order: 4,
		},
		{
			fname: "Phyllis",
			lname: "Vance",
			age:   45,
			order: 5,
		},
		{
			fname: "Kevin",
			lname: "Malone",
			age:   39,
			order: 4,
		},
		{
			fname: "Oscar",
			lname: "Martinez",
			age:   37,
			order: 5,
		},
		{
			fname: "Meredith",
			lname: "Palmer",
			age:   43,
			order: 4,
		},
		{
			fname: "Pam",
			lname: "Beesly",
			age:   30,
			order: 6,
		},
		{
			fname: "Angela",
			lname: "Martin",
			age:   33,
			order: 4,
		},
		{
			fname: "Andrew",
			lname: "Bernard",
			age:   35,
			order: 5,
		},
		{
			fname: "Ellie",
			lname: "Kemper",
			age:   23,
			order: 3,
		},
		{
			fname: "Stanley",
			lname: "Hudson",
			age:   53,
			order: 3,
		},
		{
			fname: "Darryl",
			lname: "Phylbin",
			age:   36,
			order: 2,
		},
		{
			fname: "Toby",
			lname: "Flendersen",
			age:   39,
			order: 4,
		},
		{
			fname: "Michael",
			lname: "Scott",
			age:   45,
			order: 14,
		},
	}
}

func genSortMethod() sortMethod {
	a := getTestData()
	return func(i, j int) bool {
		return a[i].fname < a[j].fname
	}
}

var result []emp

func BenchmarkHeapSort(b *testing.B) {
	var r []emp
	a := getTestData()
	less := genSortMethod()
	for i := 0; i < b.N; i++ {
		r = HeapSort(a, less)
	}
	result = r
}

func benchmarkNormalUnstableSort(b *testing.B) {
	less := genSortMethod()
	for i := 0; i < b.N; i++ {
		sort.Slice(getTestData(), less)
	}
}

func BenchmarkCompareHeapSort(b *testing.B) {
	tests := map[string]struct {
		bench func(b *testing.B)
	}{
		"HeapSort": {
			bench: BenchmarkHeapSort,
		},
		"NormalUnstableSort": {
			bench: benchmarkNormalUnstableSort,
		},
	}
	for name, test := range tests {
		b.Run(name, test.bench)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	var r []emp
	a := getTestData()
	less := genSortMethod()
	for i := 0; i < b.N; i++ {
		r = MergeSort(a, less)
	}
	result = r
}

func benchmarkNormalStableSort(b *testing.B) {
	less := genSortMethod()
	for i := 0; i < b.N; i++ {
		sort.SliceStable(getTestData(), less)
	}
}

func BenchmarkCompareMergeSort(b *testing.B) {
	tests := map[string]struct {
		bench func(b *testing.B)
	}{
		"MergeSort": {
			bench: BenchmarkMergeSort,
		},
		"NormalStableSort": {
			bench: benchmarkNormalStableSort,
		},
	}
	for name, test := range tests {
		b.Run(name, test.bench)
	}
}
