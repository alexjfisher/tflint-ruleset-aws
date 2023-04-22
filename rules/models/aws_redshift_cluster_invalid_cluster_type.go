// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftClusterInvalidClusterTypeRule checks the pattern is valid
type AwsRedshiftClusterInvalidClusterTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftClusterInvalidClusterTypeRule returns new rule with default attributes
func NewAwsRedshiftClusterInvalidClusterTypeRule() *AwsRedshiftClusterInvalidClusterTypeRule {
	return &AwsRedshiftClusterInvalidClusterTypeRule{
		resourceType:  "aws_redshift_cluster",
		attributeName: "cluster_type",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftClusterInvalidClusterTypeRule) Name() string {
	return "aws_redshift_cluster_invalid_cluster_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftClusterInvalidClusterTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftClusterInvalidClusterTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftClusterInvalidClusterTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftClusterInvalidClusterTypeRule) Check(runner tflint.Runner) error {
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
					"cluster_type must be 2147483647 characters or less",
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
