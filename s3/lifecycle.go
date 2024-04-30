package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

func (b *Bucket) CreateLifecycle(ctx context.Context, client *s3.Client, params *BucketLifecycleInput) error {

	_, err := client.PutBucketLifecycleConfiguration(ctx, &s3.PutBucketLifecycleConfigurationInput{
		Bucket: aws.String(b.BucketInput.Name),
		LifecycleConfiguration: &types.BucketLifecycleConfiguration{
			Rules: []types.LifecycleRule{
				{
					ID:     aws.String(uuid.New().String()),
					Status: types.ExpirationStatusEnabled,
					Expiration: &types.LifecycleExpiration{
						Days: aws.Int32(params.RetentionInDays),
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to set bucket lifecycle configuration, %v", err)
		return err
	}

	return nil
}
