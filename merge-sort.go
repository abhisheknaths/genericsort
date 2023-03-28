package genericsort

// Make sure `source` is closed inside less
func MergeSort[T any](source []T, less func(i, j int) bool) []T {
	l := len(source)
	if l == 0 {
		return source
	}
	left, right := 0, l-1
	result := make([]T, 0, l)
	remappedIndexList := merge_sort(left, right, less)
	for i := range remappedIndexList {
		result = append(result, source[remappedIndexList[i]])
	}
	return result
}

func merge_sort(left, right int, less func(i, j int) bool) []int {
	if right <= left {
		return []int{left}
	}
	m := (right + left) / 2
	l := merge_sort(left, m-1, less)
	if left == m {
		m = m + 1
	}
	r := merge_sort(m, right, less)
	return merge(l, r, less)
}

func merge(left, right []int, less func(i, j int) bool) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) != 0 && len(right) != 0 {
		if less(left[0], right[0]) {
			result = append(result, left[0])
			left = left[1:]
		} else if less(right[0], left[0]) {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
		}
	}
	if len(left) > 0 {
		result = append(result, left...)
	} else {
		result = append(result, right...)
	}
	return result
}
