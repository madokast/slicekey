package slicekey

import (
	"fmt"
	"reflect"
)

type Slice[E any] struct {
	data any // array obj
}

// create a slicekey from elements
func Of[E any](elements ...E) Slice[E] {
	sliceType := reflect.TypeOf(elements)
	eleType := sliceType.Elem()

	arrType := reflect.ArrayOf(len(elements), eleType) // array with same length and element-type of es

	arrObj := reflect.New(arrType).Elem() // new array obj and de-ref
	for i, e := range elements {
		arrEleObj := arrObj.Index(i)
		arrEleObj.Set(reflect.ValueOf(e))
	}

	return Slice[E]{data: arrObj.Interface()}
}

// create a slicekey from slice
func Create[E any](es []E) Slice[E] {
	return Of(es...)
}

// the length of the slicekey
func (s *Slice[E]) Len() int {
	arrObj := reflect.ValueOf(s.data)
	return arrObj.Len()
}

func (s *Slice[E]) Get(index int) E {
	arrObj := reflect.ValueOf(s.data)
	obj := arrObj.Index(index)
	return obj.Interface().(E)
}

// get a slice copied from the slicekey
func (s *Slice[E]) Slice() []E {
	es := make([]E, s.Len())
	esObj := reflect.ValueOf(es)
	arrObj := reflect.ValueOf(s.data)
	for i := 0; i < s.Len(); i++ {
		esObj.Index(i).Set(arrObj.Index(i))
	}
	return es
}

func (s *Slice[E]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}
