// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElasticBeanstalkApplication(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Elasticbeanstalkconn.DescribeApplicationsRequest(&elasticbeanstalk.DescribeApplicationsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Applications) > 0 {

		for _, r := range resp.Applications {

			result = append(result, terraform.Resource{
				Type:      "aws_elastic_beanstalk_application",
				ID:        *r.ApplicationName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
