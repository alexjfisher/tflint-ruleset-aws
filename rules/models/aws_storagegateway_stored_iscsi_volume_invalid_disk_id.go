// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule checks the pattern is valid
type AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule returns new rule with default attributes
func NewAwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule() *AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule {
	return &AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule{
		resourceType:  "aws_storagegateway_stored_iscsi_volume",
		attributeName: "disk_id",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule) Name() string {
	return "aws_storagegateway_stored_iscsi_volume_invalid_disk_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidDiskIDRule) Check(runner tflint.Runner) error {
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
					"disk_id must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"disk_id must be 1 characters or higher",
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
