package slicekey

import (
	"fmt"
	"reflect"
	"unsafe"
)

type bytes string

type Slice[E any] struct {
	typeInfo *E
	length   int
	data     bytes
}

func MakeSlice[E any]() Slice[E] {
	var typeInfo = (*E)(unsafe.Pointer(uintptr(1)))
	return Slice[E]{
		typeInfo: typeInfo,
		length:   0,
		data:     "",
	}
}

func MakeSliceWithCap[E any](cap int) Slice[E] {
	var typeInfo = (*E)(unsafe.Pointer(uintptr(1)))
	sz := int(unsafe.Sizeof(*typeInfo))
	return Slice[E]{
		typeInfo: typeInfo,
		length:   0,
		data:     newBytes(sz * cap),
	}
}

func (s *Slice[E]) Cap() int {
	return s.data.cap() / int(unsafe.Sizeof(*s.typeInfo))
}

func (s *Slice[E]) Add(e E) {
	if s.length == s.Cap() {
		nbs := newBytes(s.data.cap()*2 + int(unsafe.Sizeof(e)))
		if s.Cap() > 0 {
			memcopy(s.data.rowPointer(), 0, nbs)
		}
		s.data = nbs
	}

	memcopy(&e, int(unsafe.Sizeof(e))*s.length, s.data)
	s.length++
}

func (s *Slice[E]) Set(index int, e E) {
	if index >= s.length {
		panic(fmt.Sprintf("slice out of bound [%d,%d]", index, s.length))
	}
	memcopy(&e, int(unsafe.Sizeof(e))*index, s.data)
}

func (s *Slice[E]) ToSlice() []E {
	return unsafe.Slice((*E)(unsafe.Pointer(s.data.rowPointer())), s.length)
}

func newBytes(sz int) bytes {
	bs := make([]byte, sz)
	ptr := sliceHeader(bs).Data
	return asBytes(ptr, sz)
}

func (bs bytes) toByteSlice() []byte {
	ptr := bs.rowPointer()
	return unsafe.Slice(ptr, len(bs))
}

func (bs bytes) rowPointer() *byte {
	return (*byte)(unsafe.Pointer(stringHeader(string(bs)).Data))
}

func (bs bytes) cap() int {
	return len(bs)
}

func memcopy[E any](src *E, offset int, bs bytes) {
	srcHelper := unsafe.Slice((*byte)(unsafe.Pointer(src)), unsafe.Sizeof(*src))
	copy(bs.toByteSlice()[offset:], srcHelper)
}

func sliceHeader[E any](s []E) reflect.SliceHeader {
	return *(*reflect.SliceHeader)(unsafe.Pointer(&s))
}

func stringHeader(s string) reflect.StringHeader {
	return *(*reflect.StringHeader)(unsafe.Pointer(&s))
}

func asBytes(ptr uintptr, length int) bytes {
	sh := reflect.StringHeader{
		Data: ptr,
		Len:  length,
	}
	return *(*bytes)(unsafe.Pointer(&sh))
}
