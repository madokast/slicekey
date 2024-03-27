package slicekey

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
)

type Slice[E any] struct {
	data any // array obj
}

// Of create a slicekey from elements
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

// Create a slicekey from slice
func Create[E any](es []E) Slice[E] {
	return Of(es...)
}

// Len get the length of the slicekey
func (s *Slice[E]) Len() int {
	arrObj := reflect.ValueOf(s.data)
	return arrObj.Len()
}

// Get the element at index
func (s *Slice[E]) Get(index int) E {
	arrObj := reflect.ValueOf(s.data)
	obj := arrObj.Index(index)
	return obj.Interface().(E)
}

// Foreach iterate a slice-key
func (s *Slice[E]) Foreach(consumer func(index int, element E)) {
	arrObj := reflect.ValueOf(s.data)
	for i := 0; i < s.Len(); i++ {
		arrEleObj := arrObj.Index(i)
		consumer(i, arrEleObj.Interface().(E))
	}
}

// Set element at index on a copied slice-key
// the origin slice-key remains no changed
func (s *Slice[E]) Set(index int, value E) Slice[E] {
	eleType := reflect.TypeOf(value)
	arrType := reflect.ArrayOf(s.Len(), eleType)
	arrObj := reflect.New(arrType).Elem()
	s.Foreach(func(index int, e E) {
		arrObj.Index(index).Set(reflect.ValueOf(e))
	})
	arrObj.Index(index).Set(reflect.ValueOf(value))
	return Slice[E]{data: arrObj.Interface()}
}

// Slice return a go-slice copied from the slicekey
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

func (s Slice[E]) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(s.Slice())
	return buf.Bytes(), err
}

func (s *Slice[E]) UnmarshalJSON(data []byte) error {
	var es []E
	err := json.NewDecoder(bytes.NewBuffer(data)).Decode(&es)
	if err != nil {
		return err
	}
	temp := Create(es)
	s.data = temp.data
	return nil
}

func (s Slice[E]) MarshalBinary() (data []byte, err error) {
	buf := bytes.Buffer{}
	err = gob.NewEncoder(&buf).Encode(s.Slice())
	return buf.Bytes(), err
}

func (s *Slice[E]) UnmarshalBinary(data []byte) error {
	var es []E
	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&es)
	if err != nil {
		return err
	}
	temp := Create(es)
	s.data = temp.data
	return nil
}
