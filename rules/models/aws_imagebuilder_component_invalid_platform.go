// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderComponentInvalidPlatformRule checks the pattern is valid
type AwsImagebuilderComponentInvalidPlatformRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsImagebuilderComponentInvalidPlatformRule returns new rule with default attributes
func NewAwsImagebuilderComponentInvalidPlatformRule() *AwsImagebuilderComponentInvalidPlatformRule {
	return &AwsImagebuilderComponentInvalidPlatformRule{
		resourceType:  "aws_imagebuilder_component",
		attributeName: "platform",
		enum: []string{
			"Windows",
			"Linux",
		},
	}
}

// Name returns the rule name
func (r *AwsImagebuilderComponentInvalidPlatformRule) Name() string {
	return "aws_imagebuilder_component_invalid_platform"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderComponentInvalidPlatformRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderComponentInvalidPlatformRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderComponentInvalidPlatformRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderComponentInvalidPlatformRule) Check(runner tflint.Runner) error {
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
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as platform`, truncateLongMessage(val)),
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
