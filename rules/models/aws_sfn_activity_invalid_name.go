// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSfnActivityInvalidNameRule checks the pattern is valid
type AwsSfnActivityInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSfnActivityInvalidNameRule returns new rule with default attributes
func NewAwsSfnActivityInvalidNameRule() *AwsSfnActivityInvalidNameRule {
	return &AwsSfnActivityInvalidNameRule{
		resourceType:  "aws_sfn_activity",
		attributeName: "name",
		max:           80,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSfnActivityInvalidNameRule) Name() string {
	return "aws_sfn_activity_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSfnActivityInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSfnActivityInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSfnActivityInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSfnActivityInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 80 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
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
