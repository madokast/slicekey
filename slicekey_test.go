package slicekey

import (
	"math/rand"
	"testing"
)

func TestOfInt(t *testing.T) {
	intArr := Of(1, 2, 3)
	t.Logf("%T %v", intArr, intArr)
}

func TestOfString(t *testing.T) {
	intArr := Of("", "hello", "apple")
	t.Logf("%T %v", intArr, intArr)
}

func TestLen(t *testing.T) {
	for i := 0; i < 5; i++ {
		s := make([]int, i)
		sk := Of(s...)
		t.Log(sk.Len())
	}
}

func TestLen2(t *testing.T) {
	for i := 0; i < 100; i++ {
		le := int(rand.Int31() % 1000)
		s := make([]int, le)
		sk := Of(s...)
		if sk.Len() != le {
			t.Fatalf("len neq %d <> %d", le, sk.Len())
			t.Fail()
		}
	}
}

func TestSlice(t *testing.T) {
	intArr := Of(1, 2, 3)
	intS := intArr.Slice()
	t.Logf("%T %v", intS, intS)
}

func TestString(t *testing.T) {
	intArr := Of(1, 2, 3)
	t.Logf("%T %v", intArr, intArr)
}

func TestMapKey(t *testing.T) {
	m := map[Slice[int]]string{}
	m[Of([]int{}...)] = "empty"
	m[Of(1)] = "1-"
	m[Of(1, 2)] = "1-2"
	m[Of(1, 3)] = "1-3"
	m[Of(3, 1)] = "3-1"

	m[Of(1)] = "1-modify"

	t.Logf("%T %v", m, m)

	t.Logf("m[1, 3] = %s", m[Of(1, 3)])
	t.Logf("m[3, 1] = %s", m[Of(3, 1)])
}
