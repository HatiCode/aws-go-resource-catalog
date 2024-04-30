package s3

type Bucket struct {
	BucketInput *BucketInput
}

type BucketInput struct {
	Name            string
	RetentionPeriod int32
	StorageType     string
	Lifecycle       *BucketLifecycleInput
}

type BucketLifecycleInput struct {
	RetentionInDays int32
	StorageType     string
}
