// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAmplifyBranchInvalidBasicAuthCredentialsRule checks the pattern is valid
type AwsAmplifyBranchInvalidBasicAuthCredentialsRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsAmplifyBranchInvalidBasicAuthCredentialsRule returns new rule with default attributes
func NewAwsAmplifyBranchInvalidBasicAuthCredentialsRule() *AwsAmplifyBranchInvalidBasicAuthCredentialsRule {
	return &AwsAmplifyBranchInvalidBasicAuthCredentialsRule{
		resourceType:  "aws_amplify_branch",
		attributeName: "basic_auth_credentials",
		max:           2000,
		pattern:       regexp.MustCompile(`^(?s).*$`),
	}
}

// Name returns the rule name
func (r *AwsAmplifyBranchInvalidBasicAuthCredentialsRule) Name() string {
	return "aws_amplify_branch_invalid_basic_auth_credentials"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAmplifyBranchInvalidBasicAuthCredentialsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAmplifyBranchInvalidBasicAuthCredentialsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAmplifyBranchInvalidBasicAuthCredentialsRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAmplifyBranchInvalidBasicAuthCredentialsRule) Check(runner tflint.Runner) error {
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
					"basic_auth_credentials must be 2000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`basic_auth_credentials does not match valid pattern ^(?s).*$`,
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
