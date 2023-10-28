func remove_i(int i, a []int) []int {
	return append(a[i], a[i+1]...)
}