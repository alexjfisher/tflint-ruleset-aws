// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftClusterInvalidDatabaseNameRule checks the pattern is valid
type AwsRedshiftClusterInvalidDatabaseNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftClusterInvalidDatabaseNameRule returns new rule with default attributes
func NewAwsRedshiftClusterInvalidDatabaseNameRule() *AwsRedshiftClusterInvalidDatabaseNameRule {
	return &AwsRedshiftClusterInvalidDatabaseNameRule{
		resourceType:  "aws_redshift_cluster",
		attributeName: "database_name",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftClusterInvalidDatabaseNameRule) Name() string {
	return "aws_redshift_cluster_invalid_database_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftClusterInvalidDatabaseNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftClusterInvalidDatabaseNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftClusterInvalidDatabaseNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftClusterInvalidDatabaseNameRule) Check(runner tflint.Runner) error {
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
					"database_name must be 2147483647 characters or less",
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
