// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule checks the pattern is valid
type AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule returns new rule with default attributes
func NewAwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule() *AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule {
	return &AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule{
		resourceType:  "aws_kinesisanalyticsv2_application",
		attributeName: "service_execution_role",
		max:           2048,
		min:           1,
		pattern:       regexp.MustCompile(`^arn:.*$`),
	}
}

// Name returns the rule name
func (r *AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule) Name() string {
	return "aws_kinesisanalyticsv2_application_invalid_service_execution_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKinesisanalyticsv2ApplicationInvalidServiceExecutionRoleRule) Check(runner tflint.Runner) error {
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
					"service_execution_role must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"service_execution_role must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:.*$`),
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
