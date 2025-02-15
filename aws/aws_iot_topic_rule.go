// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListIotTopicRule(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Iotconn.ListTopicRulesRequest(&iot.ListTopicRulesInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Rules) > 0 {

		for _, r := range resp.Rules {

			result = append(result, terraform.Resource{
				Type:      "aws_iot_topic_rule",
				ID:        *r.RuleName,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
