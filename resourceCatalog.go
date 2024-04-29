package main

import (
	"github.com/HatiCode/aws-go-resource-catalog/s3"
)

type ResourceCatalog interface {
	s3.Bucket
}
