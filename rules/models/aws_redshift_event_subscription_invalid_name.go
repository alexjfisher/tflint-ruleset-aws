// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftEventSubscriptionInvalidNameRule checks the pattern is valid
type AwsRedshiftEventSubscriptionInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftEventSubscriptionInvalidNameRule returns new rule with default attributes
func NewAwsRedshiftEventSubscriptionInvalidNameRule() *AwsRedshiftEventSubscriptionInvalidNameRule {
	return &AwsRedshiftEventSubscriptionInvalidNameRule{
		resourceType:  "aws_redshift_event_subscription",
		attributeName: "name",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftEventSubscriptionInvalidNameRule) Name() string {
	return "aws_redshift_event_subscription_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftEventSubscriptionInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftEventSubscriptionInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftEventSubscriptionInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftEventSubscriptionInvalidNameRule) Check(runner tflint.Runner) error {
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
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 2147483647 characters or less",
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
