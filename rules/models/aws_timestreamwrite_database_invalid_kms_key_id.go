// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule checks the pattern is valid
type AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsTimestreamwriteDatabaseInvalidKmsKeyIDRule returns new rule with default attributes
func NewAwsTimestreamwriteDatabaseInvalidKmsKeyIDRule() *AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule {
	return &AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule{
		resourceType:  "aws_timestreamwrite_database",
		attributeName: "kms_key_id",
		max:           2048,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule) Name() string {
	return "aws_timestreamwrite_database_invalid_kms_key_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTimestreamwriteDatabaseInvalidKmsKeyIDRule) Check(runner tflint.Runner) error {
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
					"kms_key_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"kms_key_id must be 1 characters or higher",
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
