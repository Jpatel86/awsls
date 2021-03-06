// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
)

func ListNeptuneEventSubscription(client *Client) ([]Resource, error) {
	req := client.neptuneconn.DescribeEventSubscriptionsRequest(&neptune.DescribeEventSubscriptionsInput{})

	var result []Resource

	p := neptune.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSubscriptionsList {

			result = append(result, Resource{
				Type:   "aws_neptune_event_subscription",
				ID:     *r.CustSubscriptionId,
				Region: client.neptuneconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
