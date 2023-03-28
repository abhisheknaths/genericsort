# Generic Sort

## Contents
This package has two different sort functions using Heap sort and merge sort algorithms. The signature for these functions are very similar to the functions provided by the [sort package](https://pkg.go.dev/sort "Go's native sort package").
The functions [Slice](https://pkg.go.dev/sort#Slice) and [SliceStable](https://pkg.go.dev/sort#SliceStable) are the reference for these new methods and like them these functions accept *[]any* as input.

## Purpose
The package can work as a reference to some go lang features and testing practices

- Parametric Polymorphism
- Closures
- Benchmarking and comparing functions based on benchmark results
- Testing the package as if it were imported from another package
- Running tests inside a container etc...

## Operations
The make file has commands for all the benchmark tests, run them to see comparisons with native functions.

## Warning

- Unlike the native methods, the algorithms this package uses are not in place, so they will return a new slice.
- The input slice should be the same one that is closed within the *less* function.

## Example Usage

```
a := []int{9, 2, 1, 6, 3, 0, 4}
a = genericsort.HeapSort(a, func(i, j int) bool {
    return a[i] < a[j]
})
```



