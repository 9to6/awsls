// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListSesReceiptFilter(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Sesconn.ListReceiptFiltersRequest(&ses.ListReceiptFiltersInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Filters) > 0 {

		for _, r := range resp.Filters {

			result = append(result, terraform.Resource{
				Type:      "aws_ses_receipt_filter",
				ID:        *r.Name,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
