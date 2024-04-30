package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/ptr"
)

func (b *Bucket) Create(ctx context.Context, params *BucketInput) (*s3.CreateBucketOutput, error) {
	if params == nil {
		params = &BucketInput{}
	}

	// Create Config and Client
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	// Create Bucket
	bc, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(params.Name),
	})

	if err != nil {
		log.Fatalf("failed to create bucket, %v", err)
		return nil, err
	}

	if b.BucketInput.Lifecycle != nil {
		b.CreateLifecycle(ctx, client, &BucketLifecycleInput{
			BucketName: ptr.String(b.BucketInput.Name),
			RetentionInDays: b.BucketInput.Lifecycle.RetentionInDays,
			StorageType: b.BucketInput.Lifecycle.StorageType,
		})
	}
	
	return bc, nil
}
