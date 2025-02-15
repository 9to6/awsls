// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSqsQueue(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Sqsconn.ListQueuesRequest(&sqs.ListQueuesInput{})

	var result []terraform.Resource

	p := sqs.NewListQueuesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.QueueUrls {

			result = append(result, terraform.Resource{
				Type:      "aws_sqs_queue",
				ID:        r,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
