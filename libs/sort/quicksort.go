package sort

func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int, l, r int) {
	if l > r {
		return
	}
	p := partition(a, l, r)
	quickSort(a, l, p-1)
	quickSort(a, p+1, r)
}
