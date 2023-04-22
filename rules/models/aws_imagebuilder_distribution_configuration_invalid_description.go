// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderDistributionConfigurationInvalidDescriptionRule checks the pattern is valid
type AwsImagebuilderDistributionConfigurationInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderDistributionConfigurationInvalidDescriptionRule returns new rule with default attributes
func NewAwsImagebuilderDistributionConfigurationInvalidDescriptionRule() *AwsImagebuilderDistributionConfigurationInvalidDescriptionRule {
	return &AwsImagebuilderDistributionConfigurationInvalidDescriptionRule{
		resourceType:  "aws_imagebuilder_distribution_configuration",
		attributeName: "description",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderDistributionConfigurationInvalidDescriptionRule) Name() string {
	return "aws_imagebuilder_distribution_configuration_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderDistributionConfigurationInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderDistributionConfigurationInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderDistributionConfigurationInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderDistributionConfigurationInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"description must be 1 characters or higher",
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
