// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsRouteInvalidGatewayRule checks whether attribute value actually exists
type AwsRouteInvalidGatewayRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsRouteInvalidGatewayRule returns new rule with default attributes
func NewAwsRouteInvalidGatewayRule() *AwsRouteInvalidGatewayRule {
	return &AwsRouteInvalidGatewayRule{
		resourceType:  "aws_route",
		attributeName: "gateway_id",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsRouteInvalidGatewayRule) Name() string {
	return "aws_route_invalid_gateway"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRouteInvalidGatewayRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRouteInvalidGatewayRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRouteInvalidGatewayRule) Link() string {
	return ""
}

// Metadata returns the metadata about deep checking
func (r *AwsRouteInvalidGatewayRule) Metadata() interface{} {
	return map[string]bool{"deep": true}
}

// Check checks whether the attributes are included in the list retrieved by DescribeInternetGateways
func (r *AwsRouteInvalidGatewayRule) Check(rr tflint.Runner) error {
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

		if !r.dataPrepared {
			awsClient, err := runner.AwsClient(resource.Body.Attributes)
			if err != nil {
				return err
			}
			logger.Debug("invoking DescribeInternetGateways")
			r.data, err = awsClient.DescribeInternetGateways()
			if err != nil {
				err := fmt.Errorf("An error occurred while invoking DescribeInternetGateways; %w", err)
				logger.Error("%s", err)
				return err
			}
			r.dataPrepared = true
		}

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid internet gateway ID.`, val),
					attribute.Expr.Range(),
				)
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
