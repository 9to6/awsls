// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSesConfigurationSet(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Sesconn.ListConfigurationSetsRequest(&ses.ListConfigurationSetsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.ConfigurationSets) > 0 {

		for _, r := range resp.ConfigurationSets {

			result = append(result, terraform.Resource{
				Type:      "aws_ses_configuration_set",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
