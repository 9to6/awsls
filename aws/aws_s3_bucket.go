// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListS3Bucket(client *aws.Client) ([]terraform.Resource, error) {
	req := client.S3conn.ListBucketsRequest(&s3.ListBucketsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Buckets) > 0 {

		for _, r := range resp.Buckets {

			t := *r.CreationDate
			result = append(result, terraform.Resource{
				Type:      "aws_s3_bucket",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,

				CreatedAt: &t,
			})
		}
	}

	return result, nil
}
