package math

import "testing"

func Test_Sum(t *testing.T) {
	arg := []int{15, 0, 4, -6}
	sum := Sum(arg)
	want := 13
	if want != sum {
		t.Errorf("Error: Sum(%v): want %d, but got %d", arg, want, sum)
	}
}
