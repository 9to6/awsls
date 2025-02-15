// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListCodeartifactRepository(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Codeartifactconn.ListRepositoriesRequest(&codeartifact.ListRepositoriesInput{})

	var result []terraform.Resource

	p := codeartifact.NewListRepositoriesPaginator(req)
	for p.Next(context.Background()) {
		resp := p.CurrentPage()

		for _, r := range resp.Repositories {

			result = append(result, terraform.Resource{
				Type:      "aws_codeartifact_repository",
				ID:        *r.Arn,
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
