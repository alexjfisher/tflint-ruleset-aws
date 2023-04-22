// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventTargetInvalidInputRule checks the pattern is valid
type AwsCloudwatchEventTargetInvalidInputRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsCloudwatchEventTargetInvalidInputRule returns new rule with default attributes
func NewAwsCloudwatchEventTargetInvalidInputRule() *AwsCloudwatchEventTargetInvalidInputRule {
	return &AwsCloudwatchEventTargetInvalidInputRule{
		resourceType:  "aws_cloudwatch_event_target",
		attributeName: "input",
		max:           8192,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventTargetInvalidInputRule) Name() string {
	return "aws_cloudwatch_event_target_invalid_input"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventTargetInvalidInputRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventTargetInvalidInputRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventTargetInvalidInputRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventTargetInvalidInputRule) Check(runner tflint.Runner) error {
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
					"input must be 8192 characters or less",
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
