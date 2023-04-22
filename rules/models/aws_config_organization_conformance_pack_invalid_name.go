// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConfigOrganizationConformancePackInvalidNameRule checks the pattern is valid
type AwsConfigOrganizationConformancePackInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsConfigOrganizationConformancePackInvalidNameRule returns new rule with default attributes
func NewAwsConfigOrganizationConformancePackInvalidNameRule() *AwsConfigOrganizationConformancePackInvalidNameRule {
	return &AwsConfigOrganizationConformancePackInvalidNameRule{
		resourceType:  "aws_config_organization_conformance_pack",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z][-a-zA-Z0-9]*$`),
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationConformancePackInvalidNameRule) Name() string {
	return "aws_config_organization_conformance_pack_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationConformancePackInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationConformancePackInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationConformancePackInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationConformancePackInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z][-a-zA-Z0-9]*$`),
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
