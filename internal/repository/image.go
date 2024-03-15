package repository

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/constant"
)

type image struct {
	client *s3.S3
}

func NewImageRepository(client *s3.S3) internal.ImageRepository {
	return &image{
		client: client,
	}
}

// Upload implements internal.ImageRepository.
func (r *image) Upload(ctx context.Context, bucket string, filePath string, file *bytes.Buffer) error {
	_, err := r.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
		ACL:    aws.String(constant.AWSPubReadACL),
		Body:   bytes.NewReader(file.Bytes()),
	})

	return err
}
