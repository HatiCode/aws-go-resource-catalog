package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (b *Bucket) New(ctx context.Context, params *BucketInput) (*s3.CreateBucketOutput, error) {
	if params == nil {
		params = &BucketInput{}
	}

	// Create Config and Client
	cfg, err := config.LoadDefaultConfig(context.TODO())
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

	// lc, err := b.CreateLifecycle(ctx, client, &LifecycleInput{
	// 	BucketName:      params.Name,
	// 	RetentionInDays: params.Lifecycle.LifecycleInput.RetentionInDays,
	// })

	// if err != nil {
	// 	log.Fatalf("failed to set bucket lifecycle configuration, %v", err)
	// 	return nil, err
	// }

	// return &BucketOutput{
	// 	Name:     params.Name,
	// 	Metadata: bc.ResultMetadata,
	// }, nil

	return bc, nil
}
