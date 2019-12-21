package math

import "testing"

func Test_Sum(t *testing.T) {
	sum := Sum([]int{15, 0, 4, -6})
	want := 13
	if want != sum {
		t.Errorf("Error: Sum(): want %d, but got %d", want, sum)
	}
}
