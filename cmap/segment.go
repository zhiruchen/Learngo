package cmap

import (
	"sync"
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
