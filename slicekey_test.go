package slicekey

import (
	"fmt"
	"slices"
	"testing"
)

func TestMake(t *testing.T) {
	s := MakeSlice[int]()
	assert(len(s.ToSlice()) == 0, "%v", s.ToSlice())
}

func TestMakeSliceWithCap(t *testing.T) {
	s := MakeSliceWithCap[int](5)
	assert(len(s.ToSlice()) == 0, "%v", s.ToSlice())
}

func TestCap0(t *testing.T) {
	s := MakeSlice[int]()
	assert(s.Cap() == 0, "%v", s.Cap())
}

func TestCap1(t *testing.T) {
	s := MakeSliceWithCap[int](10)
	assert(s.Cap() == 10, "%v", s.Cap())
}

func TestCap2(t *testing.T) {
	s := MakeSliceWithCap[string](10)
	assert(s.Cap() == 10, "%v", s.Cap())
}

func TestAdd(t *testing.T) {
	s := MakeSlice[int]()
	s.Add(12)
	assert(slices.Equal(s.ToSlice(), []int{12}), "%v", s.ToSlice())
}


func assert(b bool, format string, args...any) {
	if !b {
		panic(fmt.Sprintf(format, args...))
	}
}