// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListEc2CapacityReservation(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Ec2conn.DescribeCapacityReservationsRequest(&ec2.DescribeCapacityReservationsInput{})

	var result []terraform.Resource

	p := ec2.NewDescribeCapacityReservationsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.CapacityReservations {
			if *r.OwnerId != client.AccountID {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreateDate
			result = append(result, terraform.Resource{
				Type:      "aws_ec2_capacity_reservation",
				ID:        *r.CapacityReservationId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
