package genericsort

// Make sure `source` is closed inside less
func HeapSort[T any](source []T, less func(i, j int) bool) []T {
	result := make([]T, 0, len(source))
	heapify(source, less)
	var lastIndex int
	for len(source) > 0 {
		lastIndex = len(source) - 1
		result = append(result, source[0])
		source[0], source[lastIndex] = source[lastIndex], source[0]
		source = source[:lastIndex]
		var siftDownIndex int
		var ok bool
		for {
			siftDownIndex, ok = sift_down(source, siftDownIndex, less)
			if !ok {
				break
			}
		}
	}
	return result
}

func heapify[T any](source []T, less func(i, j int) bool) {
	l := len(source)
	for i := l - 1; i >= 0; i-- {
		k := i
		var ok bool
		for {
			k, ok = sift_down(source, k, less)
			if !ok {
				break
			}
		}
	}
}

func sift_down[T any](source []T, index int, less func(i, j int) bool) (int, bool) {
	l := len(source)
	child1 := (2 * index) + 1
	if child1 >= l {
		return index, false
	}
	child2 := child1 + 1
	if child2 >= l {
		if less(child1, index) {
			source[index], source[child1] = source[child1], source[index]
			return child1, true
		}
		return index, false
	}
	var candidateIndex int
	if less(child1, child2) {
		candidateIndex = child1
	} else {
		candidateIndex = child2
	}
	if less(candidateIndex, index) {
		source[index], source[candidateIndex] = source[candidateIndex], source[index]
		return candidateIndex, true
	}
	return index, false
}
