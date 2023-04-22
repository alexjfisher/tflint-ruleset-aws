// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCustomerGatewayInvalidTypeRule checks the pattern is valid
type AwsCustomerGatewayInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCustomerGatewayInvalidTypeRule returns new rule with default attributes
func NewAwsCustomerGatewayInvalidTypeRule() *AwsCustomerGatewayInvalidTypeRule {
	return &AwsCustomerGatewayInvalidTypeRule{
		resourceType:  "aws_customer_gateway",
		attributeName: "type",
		enum: []string{
			"ipsec.1",
		},
	}
}

// Name returns the rule name
func (r *AwsCustomerGatewayInvalidTypeRule) Name() string {
	return "aws_customer_gateway_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCustomerGatewayInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCustomerGatewayInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCustomerGatewayInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCustomerGatewayInvalidTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
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

		err := runner.EvaluateExpr(attribute.Expr, func (val string) error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
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
