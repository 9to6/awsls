// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListElastictranscoderPreset(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Elastictranscoderconn.ListPresetsRequest(&elastictranscoder.ListPresetsInput{})

	var result []terraform.Resource

	p := elastictranscoder.NewListPresetsPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Presets {

			result = append(result, terraform.Resource{
				Type:      "aws_elastictranscoder_preset",
				ID:        *r.Id,
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
