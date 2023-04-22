// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule checks the pattern is valid
type AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule returns new rule with default attributes
func NewAwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule() *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule {
	return &AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule{
		resourceType:  "aws_redshift_snapshot_schedule_association",
		attributeName: "cluster_identifier",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Name() string {
	return "aws_redshift_snapshot_schedule_association_invalid_cluster_identifier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftSnapshotScheduleAssociationInvalidClusterIdentifierRule) Check(runner tflint.Runner) error {
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
					"cluster_identifier must be 2147483647 characters or less",
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
