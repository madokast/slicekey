package slicekey

import (
	"fmt"
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

	key := make([]int, 0)
	m[Of(key...)] = "empty"

	key1 := make([]int, 1)
	key1[0] = 1
	m[Of(key1...)] = "1-"

	key2 := make([]int, 2)
	key2[0] = 1
	key2[1] = 3
	m[Of(key2...)] = "1-3"

	key2[0] = 3
	key2[1] = 1
	m[Of(key2...)] = "3-1"

	m[Of(1)] = "1-modify"

	fmt.Printf("%v\n", m) // map[{[1]}:1-modify {[]}:empty {[1 3]}:1-3 {[3 1]}:3-1]

	fmt.Printf("m[1, 3] = %s\n", m[Of(1, 3)]) // m[1, 3] = 1-3
	fmt.Printf("m[3, 1] = %s\n", m[Of(3, 1)]) // m[3, 1] = 3-1
}

func TestMapKey2(t *testing.T) {
	m := make(map[Slice[int]]int)

	for i := 100; i < 102; i++ {
		m[Of(i)] = i
		for j := 20; j < 23; j++ {
			m[Of(i, j)] = i + j
		}
	}

	fmt.Printf("%v\n", m) // map[{[100]}:100 {[101]}:101 {[100 20]}:120 {[100 21]}:121 {[100 22]}:122 {[101 20]}:121 {[101 21]}:122 {[101 22]}:123]
}
