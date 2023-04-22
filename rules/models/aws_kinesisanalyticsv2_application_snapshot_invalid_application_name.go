// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule checks the pattern is valid
type AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule returns new rule with default attributes
func NewAwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule() *AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule {
	return &AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule{
		resourceType:  "aws_kinesisanalyticsv2_application_snapshot",
		attributeName: "application_name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule) Name() string {
	return "aws_kinesisanalyticsv2_application_snapshot_invalid_application_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsKinesisanalyticsv2ApplicationSnapshotInvalidApplicationNameRule) Check(runner tflint.Runner) error {
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
					"application_name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"application_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9_.-]+$`),
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
