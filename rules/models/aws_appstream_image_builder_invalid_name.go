// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppstreamImageBuilderInvalidNameRule checks the pattern is valid
type AwsAppstreamImageBuilderInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAppstreamImageBuilderInvalidNameRule returns new rule with default attributes
func NewAwsAppstreamImageBuilderInvalidNameRule() *AwsAppstreamImageBuilderInvalidNameRule {
	return &AwsAppstreamImageBuilderInvalidNameRule{
		resourceType:  "aws_appstream_image_builder",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`),
	}
}

// Name returns the rule name
func (r *AwsAppstreamImageBuilderInvalidNameRule) Name() string {
	return "aws_appstream_image_builder_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppstreamImageBuilderInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppstreamImageBuilderInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppstreamImageBuilderInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppstreamImageBuilderInvalidNameRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9][a-zA-Z0-9_.-]{0,100}$`),
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
