package sort

func InsertionSort(list []int) []int {
	for j := 1; j < len(list); j++ {
		key := list[j]
		i := j - 1
		for i >= 0 && list[i] > key {
			list[i+1] = list[i]
			i--
		}
		list[i+1] = key
	}
	return list
}
