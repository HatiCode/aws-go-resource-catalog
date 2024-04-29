package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (b *Bucket) CreateLifecycle(ctx context.Context, client *s3.Client, params *LifecycleInput) (*LifecycleOutput, error) {
	input := &s3.PutBucketLifecycleConfigurationInput{
		Bucket: aws.String(params.BucketName),
		LifecycleConfiguration: &types.BucketLifecycleConfiguration{
			Rules: []types.LifecycleRule{
				{
					ID:     aws.String("rule1"),
					Status: types.ExpirationStatusEnabled,
					Expiration: &types.LifecycleExpiration{
						Days: aws.Int32(params.RetentionInDays),
					},
				},
			},
		},
	}

	apply, err := client.PutBucketLifecycleConfiguration(ctx, input)
	if err != nil {
		log.Fatalf("failed to set bucket lifecycle configuration, %v", err)
		return nil, err
	}

	return &LifecycleOutput{
		Metadata: apply,
	}, nil
}
