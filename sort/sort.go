package sort

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

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

func MergeSort(list []int, p int, r int) []int {
	if p < r {
		q := (p + r) / 2
		MergeSort(list, p, q)
		MergeSort(list, q+1, r)
		merge(list, p, q, r)
	}
	return list
}

func merge(list []int, p int, q int, r int) {
	n1 := q - p + 1
	n2 := r - q
	var L []int
	for i := 0; i < n1; i++ {
		L = append(L, list[p+i-1])
	}
	var R []int
	for i := 0; i < n2; i++ {
		R = append(R, list[q+i])
	}
	L = append(L, maxInt)
	R = append(R, maxInt)
	i := 0
	j := 0
	for k := p - 1; k < r; k++ {
		if L[i] < R[j] {
			list[k] = L[i]
			i++
		} else {
			list[k] = R[j]
			j++
		}
	}
}
