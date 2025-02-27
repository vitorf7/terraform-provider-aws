//go:build !generate
// +build !generate

package inspector

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/inspector/inspectoriface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// Custom Inspector tag service update functions using the same format as generated code.

// updateTags updates inspector resource tags.
// The identifier is the resource ARN.
func updateTags(ctx context.Context, conn inspectoriface.InspectorAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if len(newTags) > 0 {
		input := &inspector.SetTagsForResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(newTags.IgnoreAWS()),
		}

		_, err := conn.SetTagsForResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	} else if len(oldTags) > 0 {
		input := &inspector.SetTagsForResourceInput{
			ResourceArn: aws.String(identifier),
		}

		_, err := conn.SetTagsForResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates inspector service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return updateTags(ctx, meta.(*conns.AWSClient).InspectorConn(), identifier, oldTags, newTags)
}
