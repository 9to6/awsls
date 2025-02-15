// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafSqlInjectionMatchSet(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Wafconn.ListSqlInjectionMatchSetsRequest(&waf.ListSqlInjectionMatchSetsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.SqlInjectionMatchSets) > 0 {

		for _, r := range resp.SqlInjectionMatchSets {

			result = append(result, terraform.Resource{
				Type:      "aws_waf_sql_injection_match_set",
				ID:        *r.SqlInjectionMatchSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
