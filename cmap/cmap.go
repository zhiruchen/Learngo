package cmap

import (
	"errors"
	"math"
	"sync/atomic"
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

func (cmap *concurrentMap) findSegment(keyHash uint64) Segment {
	if cmap.concurrency == 1 {
		return cmap.segments[0]
	}

	var keyHash32 uint32
	if keyHash > math.MaxUint32 {
		keyHash32 = uint32(keyHash >> 32)
	} else {
		keyHash32 = uint32(keyHash)
	}

	return cmap.segments[int(keyHash32>>16)%(cmap.concurrency-1)]
}

func (cmap *concurrentMap) Get(key string) interface{} {
	keyHash := hash(key)

	s := cmap.findSegment(keyHash)
	pair := s.GetWithHash(key, keyHash)
	if pair == null {
		return null
	}

	return pair.Element()
}

func (cmap *concurrentMap) Delete(key string) bool {
	s := cmap.findSegment(hash(key))
	if s.Delete(key) {
		atomic.AddUint64(&cmap.total, ^uint64(0))
		return true
	}
	return false
}
