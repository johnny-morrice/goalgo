package quicksort

func Sort(list []int) {
	Quicksort(list, 0, len(list)-1)
}

func Quicksort(list []int, p, r int) {
	if p < r {
		q := partition(list, p, r)
		Quicksort(list, p, q-1)
		Quicksort(list, q+1, r)
	}
}

func partition(list []int, p, r int) int {
	x := list[r]
	i := p - 1
	for j := p; j < r; j++ {
		if list[j] <= x {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[r] = list[r], list[i+1]
	return i + 1
}
