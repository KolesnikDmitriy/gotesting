package math

// Sum function return sum of array elements
func Sum(a []int) int {
	var s int
	for _, el := range a {
		s += el
	}
	return s
}
