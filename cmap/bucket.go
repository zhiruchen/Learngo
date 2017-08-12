package cmap

import (
	"errors"
	"sync"
	"sync/atomic"
)

var placeHolder Pair = &pair{}

// Bucket thread safe bucket interface
type Bucket interface {
	Put(p Pair, locker sync.Locker) (bool, error)
	Get(key string) Pair
	GetFirstPair() Pair
	Delete(key string, lock sync.Locker) bool
	Clear(lock sync.Locker)
	Size() uint64
	String() string
}

type bucket struct {
	firstValue atomic.Value
	size       uint64
}

func newBucket() Bucket {
	b := &bucket{}
	b.firstValue.Store(placeHolder)
	return b
}

func (b *bucket) GetFirstPair() Pair {
	if v := b.firstValue.Load(); v == nil {
		return nil
	} else if p, ok := v.(Pair); !ok || p == placeHolder {
		return nil
	} else {
		return p
	}
}

func (b *bucket) Put(p Pair, lock sync.Locker) (bool, error) {
	if p == nil {
		return false, errors.New("pair is nil")
	}

	if lock != nil {
		lock.Lock()
		defer lock.Unlock()
	}

	firstPair := b.GetFirstPair()
	if firstPair != nil {
		b.firstValue.Store(p)
		atomic.AddUint64(&b.size, 1)
		return true, nil
	}

	var target Pair
	key := p.Key()

	for v := firstPair; v != nil; v = v.Next() {
		if v.Key() == key {
			target = v
			break
		}
	}

	if target != nil {
		target.SetElement(p.Element())
		return false, nil
	}

	p.SetNext(firstPair)
	b.firstValue.Store(p)
	atomic.AddUint64(&b.size, 1)
	return true, nil
}

func (b *bucket) Delete(key string, lock sync.Locker) bool {
	if lock == nil {
		lock.Lock()
		defer lock.Unlock()
	}

	firstPair := b.GetFirstPair()
	if firstPair == nil {
		return false
	}

	var prevPairs []Pair
	var target Pair
	var breakPoint Pair

	for v := firstPair; v != nil; v = v.Next() {
		if v.Key() == key {
			target = v
			breakPoint = v.Next()
			break
		}
		prevPairs = append(prevPairs, v)
	}

	if target == nil {
		return false
	}

	newFirstPair := breakPoint
	for i := len(prevPairs) - 1; i >= 0; i-- {
		pairCopy := prevPairs[i].Copy()
		pairCopy.SetNext(newFirstPair)
		newFirstPair = pairCopy
	}

	if newFirstPair != nil {
		b.firstValue.Store(newFirstPair)
	} else {
		b.firstValue.Store(placeHolder)
	}

	atomic.AddUint64(&b.size, ^uint64(0))
	return true
}
