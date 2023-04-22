// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEcrLifecyclePolicyInvalidRepositoryRule checks the pattern is valid
type AwsEcrLifecyclePolicyInvalidRepositoryRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsEcrLifecyclePolicyInvalidRepositoryRule returns new rule with default attributes
func NewAwsEcrLifecyclePolicyInvalidRepositoryRule() *AwsEcrLifecyclePolicyInvalidRepositoryRule {
	return &AwsEcrLifecyclePolicyInvalidRepositoryRule{
		resourceType:  "aws_ecr_lifecycle_policy",
		attributeName: "repository",
		max:           256,
		min:           2,
		pattern:       regexp.MustCompile(`^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
	}
}

// Name returns the rule name
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Name() string {
	return "aws_ecr_lifecycle_policy_invalid_repository"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrLifecyclePolicyInvalidRepositoryRule) Check(runner tflint.Runner) error {
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
					"repository must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"repository must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*$`),
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
