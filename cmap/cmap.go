package cmap

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

const MAX_CONCURRENCY = 100
const DEFAULT_BUCKET_NUMBER = 16

// ConcurrentMap concurrent goroutine safe map
type ConcurrentMap interface {
	Concurrency() int
	Put(key string, element interface{}) (bool, error)
	Get(key string) interface{}
	Delete(key string) bool
	Len() uint64
}

type concurrentMap struct {
	concurrency int
	segments    []Segment
	total       uint64 // map size
}

type Segment struct {
}

func NewConcurrentMap(concurrency int, pr PairRedistributor) (ConcurrentMap, error) {
	if concurrency <= 0 {
		return nil, errors.New("concurrency is to small")
	}

	if concurrency > MAX_CONCURRENCY {
		return nil, errors.New("concurency is too large")
	}

	cmap := &concurrentMap{}
	cmap.concurrency = concurrency
	cmap.segments = make([]Segment, concurrency)
	for i := 0; i < concurrency; i++ {
		cmap.segments[i] = newSegment(DEFAULT_BUCKET_NUMBER, pr)
	}
	return cmap, nil
}

func (cmap *concurrentMap) Put(key string, element interface{}) (bool, error) {
	p, err := newPair(key, element)
	if err != nil {
		return false, err
	}

	s := cmap.findSegment(p.Hash())
	ok, err := s.Put(p)

	if ok {
		atomic.AddUint64(&cmap.total, 1)
	}
	return ok, err
}

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

func (cmap *concurrentMap) findSegment(keyHash uint64) Segment {

}
