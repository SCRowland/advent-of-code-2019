package sliceutils

import "testing"

var equalSlices = []struct {
	l []int
	r []int
}{
	{[]int{12, 0, 0, 0, 99}, []int{12, 0, 0, 0, 99}},
	{[]int{12}, []int{12}},
	{[]int{}, []int{}},
	{nil, nil},
}

func TestTestEqTrue(t *testing.T) {
	for _, tt := range equalSlices {
		got := TestIntSliceEq(tt.l, tt.r)
		if got != true {
			t.Errorf("testIntSliceEq(%+v, %+v) = %t, want %t", tt.l, tt.r, got, false)
		}
	}
}

var nonEqualSlices = []struct {
	l []int
	r []int
}{
	{[]int{12, 0, 0, 0, 99}, []int{4, 0, 0, 0, 99}},
	{[]int{12, 0, 0, 0, 99}, []int{4, 0}},
	{nil, []int{12}},
	{[]int{12}, nil},
}

func TestTestEqFalse(t *testing.T) {
	for _, tt := range nonEqualSlices {
		got := TestIntSliceEq(tt.l, tt.r)
		if got != false {
			t.Errorf("testIntSliceEq(%+v, %+v) = %t, want %t", tt.l, tt.r, got, true)
		}
	}
}
