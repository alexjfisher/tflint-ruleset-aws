package api

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/smithy-go"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsLaunchConfigurationInvalidImageIDRule checks whether "aws_instance" has invalid AMI ID
type AwsLaunchConfigurationInvalidImageIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	amiIDs        map[string]bool
}

// NewAwsLaunchConfigurationInvalidImageIDRule returns new rule with default attributes
func NewAwsLaunchConfigurationInvalidImageIDRule() *AwsLaunchConfigurationInvalidImageIDRule {
	return &AwsLaunchConfigurationInvalidImageIDRule{
		resourceType:  "aws_launch_configuration",
		attributeName: "image_id",
		amiIDs:        map[string]bool{},
	}
}

// Name returns the rule name
func (r *AwsLaunchConfigurationInvalidImageIDRule) Name() string {
	return "aws_launch_configuration_invalid_image_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLaunchConfigurationInvalidImageIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLaunchConfigurationInvalidImageIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLaunchConfigurationInvalidImageIDRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsLaunchConfigurationInvalidImageIDRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether "aws_instance" has invalid AMI ID
func (r *AwsLaunchConfigurationInvalidImageIDRule) Check(rr tflint.Runner) error {
	runner := rr.(*aws.Runner)

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
			{Name: "provider"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		awsClient, err := runner.AwsClient(resource.Body.Attributes)
		if err != nil {
			return err
		}

		err = runner.EvaluateExpr(attribute.Expr, func(ami string) error {
			if !r.amiIDs[ami] {
				logger.Debug("Fetch AMI images: %s", ami)
				resp, err := awsClient.DescribeImages(&ec2.DescribeImagesInput{
					ImageIds: []string{ami},
				})
				if err != nil {
					var aerr smithy.APIError
					if errors.As(err, &aerr) {
						switch aerr.ErrorCode() {
						case "InvalidAMIID.Malformed":
							fallthrough
						case "InvalidAMIID.NotFound":
							fallthrough
						case "InvalidAMIID.Unavailable":
							runner.EmitIssue(
								r,
								fmt.Sprintf("\"%s\" is invalid image ID.", ami),
								attribute.Expr.Range(),
							)
							return nil
						}
					}
					err := fmt.Errorf("An error occurred while describing images; %w", err)
					logger.Error("%s", err)
					return err
				}

				if len(resp) != 0 {
					for imageID, exists := range resp {
						r.amiIDs[imageID] = exists
					}
				} else {
					runner.EmitIssue(
						r,
						fmt.Sprintf("\"%s\" is invalid image ID.", ami),
						attribute.Expr.Range(),
					)
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
