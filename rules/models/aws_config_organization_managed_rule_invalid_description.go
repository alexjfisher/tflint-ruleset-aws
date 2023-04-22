// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationManagedRuleInvalidDescriptionRule checks the pattern is valid
type AwsConfigOrganizationManagedRuleInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsConfigOrganizationManagedRuleInvalidDescriptionRule returns new rule with default attributes
func NewAwsConfigOrganizationManagedRuleInvalidDescriptionRule() *AwsConfigOrganizationManagedRuleInvalidDescriptionRule {
	return &AwsConfigOrganizationManagedRuleInvalidDescriptionRule{
		resourceType:  "aws_config_organization_managed_rule",
		attributeName: "description",
		max:           256,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationManagedRuleInvalidDescriptionRule) Name() string {
	return "aws_config_organization_managed_rule_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationManagedRuleInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationManagedRuleInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationManagedRuleInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationManagedRuleInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
					"description must be 256 characters or less",
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
