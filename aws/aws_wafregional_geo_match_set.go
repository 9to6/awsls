// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListWafregionalGeoMatchSet(client *aws.Client) ([]terraform.Resource, error) {
	req := client.Wafregionalconn.ListGeoMatchSetsRequest(&wafregional.ListGeoMatchSetsInput{})

	var result []terraform.Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.GeoMatchSets) > 0 {

		for _, r := range resp.GeoMatchSets {

			result = append(result, terraform.Resource{
				Type:      "aws_wafregional_geo_match_set",
				ID:        *r.GeoMatchSetId,
				Profile:   client.Profile,
				Region:    client.Region,
				AccountID: client.AccountID,
			})
		}
	}

	return result, nil
}
