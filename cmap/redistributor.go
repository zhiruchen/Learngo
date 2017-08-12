package cmap

type PairRedistributor interface {
	UpdateThreshold(pairCount uint64, bucketNumber int)
	CheckBucketStatus(pairCount uint64, bucketSize uint64) (bucketStatus BucketStatus)
	Redistribe(bucketStatus BucketStatus, buckets []Bucket) (newBuckets []Bucket, changed bool)
}
