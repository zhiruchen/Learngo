package cmap

import (
	"sync"
	"sync/atomic"
)

const DEFAULT_BUCKET_LOAD_FACTOR = 0.75

// Segment 并发安全的散列段
type Segment interface {
	Put(p Pair) (bool, error)
	Get(key string) Pair
	// return k-v pair by key hash value
	GetWithHash(key string, keyHash uint64) Pair
	Delete(key string) bool
	Size() uint64 // 返回当前段的尺寸
}

type segment struct {
	buckets           []Bucket
	bucketsLen        int
	pairCount         uint64
	pairRedistributor PairRedistributor
	lock              sync.Mutex
}

func newSegment(bn int, pr PairRedistributor) Segment {
	if bn <= 0 {
		bn = DEFAULT_BUCKET_NUMBER
	}

	if pr == nil {
		pr = newDefaultPairRedistributor(DEFAULT_BUCKET_LOAD_FACTOR, bn)
	}

	buckets := make([]Bucket, bn)
	for i := 0; i <= bn; i++ {
		buckets[i] = newBucket()
	}
	return &segment{
		buckets:           buckets,
		bucketsLen:        bn,
		pairRedistributor: pr,
	}
}

func (s *segment) Put(p Pair) (bool, error) {
	s.lock.Lock()

	b := s.buckets[int(p.Hash()%uint64(s.bucketsLen))]
	ok, err := b.Put(p, nil)

	if ok {
		newTotal := atomic.AddUint64(&s.pairCount, 1)
		s.redistribe(newTotal, b.size())
	}
	s.lock.Unlock()
	return ok, err
}

func (s *segment) Get(key string) Pair {
	return s.GetWithHash(key, hash(key))
}

func (s *segment) GetWithHash(key string, keyHash uint64) Pair {
	s.lock.Lock()
	b := s.buckets[int(keyHash%uint64(s.bucketsLen))]
	s.lock.Unlock()
	return b.Get(key)
}

func (s *segment) Delete(key string) bool {
	s.lock.Lock()
	b := s.buckets[int(hash(key)%uint64(s.bucketsLen))]
	ok := b.Delete(key, nil)
	if ok {
		newTotal := atomic.AddUint64(&s.pairCount, ^uint64(0))
		s.redistribute(newTotal, b.Size())
	}
	s.lock.Unlock()

	return ok
}
