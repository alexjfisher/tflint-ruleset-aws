// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsFsxBackupInvalidFileSystemIDRule checks the pattern is valid
type AwsFsxBackupInvalidFileSystemIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFsxBackupInvalidFileSystemIDRule returns new rule with default attributes
func NewAwsFsxBackupInvalidFileSystemIDRule() *AwsFsxBackupInvalidFileSystemIDRule {
	return &AwsFsxBackupInvalidFileSystemIDRule{
		resourceType:  "aws_fsx_backup",
		attributeName: "file_system_id",
		max:           21,
		min:           11,
		pattern:       regexp.MustCompile(`^(fs-[0-9a-f]{8,})$`),
	}
}

// Name returns the rule name
func (r *AwsFsxBackupInvalidFileSystemIDRule) Name() string {
	return "aws_fsx_backup_invalid_file_system_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFsxBackupInvalidFileSystemIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFsxBackupInvalidFileSystemIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFsxBackupInvalidFileSystemIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFsxBackupInvalidFileSystemIDRule) Check(runner tflint.Runner) error {
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
					"file_system_id must be 21 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"file_system_id must be 11 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(fs-[0-9a-f]{8,})$`),
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
