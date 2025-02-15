// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafv2WebAclLoggingConfiguration(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Wafv2conn.ListLoggingConfigurationsRequest(&wafv2.ListLoggingConfigurationsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.LoggingConfigurations) > 0 {

		for _, r := range resp.LoggingConfigurations {

			result = append(result, terraform.Resource{
				Type:      "aws_wafv2_web_acl_logging_configuration",
				ID:        *r.ResourceArn,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
