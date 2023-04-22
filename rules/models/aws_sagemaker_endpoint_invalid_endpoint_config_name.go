// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerEndpointInvalidEndpointConfigNameRule checks the pattern is valid
type AwsSagemakerEndpointInvalidEndpointConfigNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsSagemakerEndpointInvalidEndpointConfigNameRule returns new rule with default attributes
func NewAwsSagemakerEndpointInvalidEndpointConfigNameRule() *AwsSagemakerEndpointInvalidEndpointConfigNameRule {
	return &AwsSagemakerEndpointInvalidEndpointConfigNameRule{
		resourceType:  "aws_sagemaker_endpoint",
		attributeName: "endpoint_config_name",
		max:           63,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`),
	}
}

// Name returns the rule name
func (r *AwsSagemakerEndpointInvalidEndpointConfigNameRule) Name() string {
	return "aws_sagemaker_endpoint_invalid_endpoint_config_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerEndpointInvalidEndpointConfigNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerEndpointInvalidEndpointConfigNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerEndpointInvalidEndpointConfigNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerEndpointInvalidEndpointConfigNameRule) Check(runner tflint.Runner) error {
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
					"endpoint_config_name must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`),
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
