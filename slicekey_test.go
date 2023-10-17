package slicekey

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
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

func TestGet(t *testing.T) {
	s := Of(1, 2, 3)
	for i := 0; i < s.Len(); i++ {
		fmt.Printf("%d\t", s.Get(i))
		if s.Get(i) != s.Slice()[i] {
			t.Fatalf("s.Get(%d)[%d] != s.Slice()[%d][%d]", i, s.Get(i), i, s.Slice()[i])
			t.Fail()
		}
	}
	fmt.Println()
}

func TestGet2(t *testing.T) {
	s := Of("abc", "hello", " world!")
	for i := 0; i < s.Len(); i++ {
		fmt.Printf("%s\t", s.Get(i))
		if s.Get(i) != s.Slice()[i] {
			t.Fatalf("s.Get(%d)[%s] != s.Slice()[%d][%s]", i, s.Get(i), i, s.Slice()[i])
			t.Fail()
		}
	}
	fmt.Println()
}

func TestEmpty(t *testing.T) {
	s := Of[int]()
	t.Log(s.String())

	s = Create[int](nil)
	t.Log(s.String())
}

func TestJson(t *testing.T) {
	s := Of("abc", "hello", " world!")
	t.Log(s.String())

	var buf = bytes.Buffer{}
	_ = json.NewEncoder(&buf).Encode(s)
	t.Log(buf.String())

	var s2 Slice[string]
	_ = json.NewDecoder(&buf).Decode(&s2)
	t.Log(s2.String())

	if s != s2 {
		panic(s.String() + " " + s2.String())
	}
}

func TestJsonEmpty(t *testing.T) {
	s := Create[int](nil)
	t.Log(s.String())

	var buf = bytes.Buffer{}
	_ = json.NewEncoder(&buf).Encode(s)
	t.Log(buf.String())

	var s2 Slice[int]
	_ = json.NewDecoder(&buf).Decode(&s2)
	t.Log(s2.String())

	if s2.Len() != 0 {
		panic(s2.Len())
	}
}

func TestGob(t *testing.T) {
	s := Of("abc", "hello", " world!")
	t.Log(s.String())

	var buf = bytes.Buffer{}
	_ = gob.NewEncoder(&buf).Encode(s)
	t.Log(buf.String())

	var s2 Slice[string]
	_ = gob.NewDecoder(&buf).Decode(&s2)
	t.Log(s2.String())

	if s != s2 {
		panic(s.String() + " " + s2.String())
	}
}

func TestGobEmpty(t *testing.T) {
	s := Create[int64](nil)
	t.Log(s.String())

	var buf = bytes.Buffer{}
	_ = gob.NewEncoder(&buf).Encode(s)
	t.Log(buf.String())

	var s2 Slice[int64]
	_ = gob.NewDecoder(&buf).Decode(&s2)
	t.Log(s2.String())

	if s != s2 {
		panic(s.String() + " " + s2.String())
	}
}
