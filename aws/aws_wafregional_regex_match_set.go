// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafregionalRegexMatchSet(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Wafregionalconn.ListRegexMatchSetsRequest(&wafregional.ListRegexMatchSetsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.RegexMatchSets) > 0 {

		for _, r := range resp.RegexMatchSets {

			result = append(result, terraform.Resource{
				Type:      "aws_wafregional_regex_match_set",
				ID:        *r.RegexMatchSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
