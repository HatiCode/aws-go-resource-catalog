package s3

import "github.com/aws/aws-sdk-go-v2/service/s3"

type Bucket struct {
	BucketInput  BucketInput
	BucketOutput BucketOutput
}

type BucketInput struct {
	Name      string
	Lifecycle BucketLifecycle
}

type BucketOutput struct {
	Name string
	Metadata *s3.CreateBucketOutput
}

type BucketLifecycle struct {
	LifecycleInput  LifecycleInput
	LifecycleOutput LifecycleOutput
}

type LifecycleInput struct {
	BucketName      string
	RetentionInDays int32
}

type LifecycleOutput struct {
	Metadata *s3.PutBucketLifecycleConfigurationOutput
}
