package cmap

import (
	"errors"
	"unsafe"
)

type Pair interface {
	linkedPair
	Key() string
	Hash() uint64
	Element() interface{}
	SetElement(element interface{}) error
	Copy() Pair
	String() string
}

type linkedPair interface {
	Next() Pair
	SetNext(nextPair Pair) error
}

type pair struct {
	key     string
	hash    uint64
	element unsafe.Pointer
	next    unsafe.Pointer
}

func newPair(key string, element interface{}) (Pair, error) {
	p := &pair{
		key:  key,
		hash: hash(key),
	}
	if element == nil {
		return nil, errors.New("element is nil")
	}

	p.element = unsafe.Pointer(&element)
	return p, nil
}
