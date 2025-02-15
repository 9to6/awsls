// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListMediaConvertQueue(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Mediaconvertconn.ListQueuesRequest(&mediaconvert.ListQueuesInput{})

	var result []terraform.Resource

	p := mediaconvert.NewListQueuesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Queues {

			result = append(result, terraform.Resource{
				Type:      "aws_media_convert_queue",
				ID:        *r.Name,
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
