package sliceutils

import "testing"

var equalSlices64 = []struct {
	l []int64
	r []int64
}{
	{[]int64{12, 0, 0, 0, 99}, []int64{12, 0, 0, 0, 99}},
	{[]int64{12}, []int64{12}},
	{[]int64{}, []int64{}},
	{nil, nil},
}

func TestTestEqTrue64(t *testing.T) {
	for _, tt := range equalSlices64 {
		got := TestIntSliceEq64(tt.l, tt.r)
		if got != true {
			t.Errorf("testIntSliceEq(%+v, %+v) = %t, want %t", tt.l, tt.r, got, false)
		}
	}
}

var nonEqualSlices64 = []struct {
	l []int64
	r []int64
}{
	{[]int64{12, 0, 0, 0, 99}, []int64{4, 0, 0, 0, 99}},
	{[]int64{12, 0, 0, 0, 99}, []int64{4, 0}},
	{nil, []int64{12}},
	{[]int64{12}, nil},
}

func TestTestEqFalse64(t *testing.T) {
	for _, tt := range nonEqualSlices64 {
		got := TestIntSliceEq64(tt.l, tt.r)
		if got != false {
			t.Errorf("testIntSliceEq(%+v, %+v) = %t, want %t", tt.l, tt.r, got, true)
		}
	}
}

var equalSlices32 = []struct {
	l []int
	r []int
}{
	{[]int{12, 0, 0, 0, 99}, []int{12, 0, 0, 0, 99}},
	{[]int{12}, []int{12}},
	{[]int{}, []int{}},
	{nil, nil},
}

func TestTestEqTrue32(t *testing.T) {
	for _, tt := range equalSlices32 {
		got := TestIntSliceEq32(tt.l, tt.r)
		if got != true {
			t.Errorf("testIntSliceEq(%+v, %+v) = %t, want %t", tt.l, tt.r, got, false)
		}
	}
}

var nonEqualSlices32 = []struct {
	l []int
	r []int
}{
	{[]int{12, 0, 0, 0, 99}, []int{4, 0, 0, 0, 99}},
	{[]int{12, 0, 0, 0, 99}, []int{4, 0}},
	{nil, []int{12}},
	{[]int{12}, nil},
}

func TestTestEqFalse32(t *testing.T) {
	for _, tt := range nonEqualSlices32 {
		got := TestIntSliceEq32(tt.l, tt.r)
		if got != false {
			t.Errorf("testIntSliceEq(%+v, %+v) = %t, want %t", tt.l, tt.r, got, true)
		}
	}
}
