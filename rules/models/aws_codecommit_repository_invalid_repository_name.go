// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodecommitRepositoryInvalidRepositoryNameRule checks the pattern is valid
type AwsCodecommitRepositoryInvalidRepositoryNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCodecommitRepositoryInvalidRepositoryNameRule returns new rule with default attributes
func NewAwsCodecommitRepositoryInvalidRepositoryNameRule() *AwsCodecommitRepositoryInvalidRepositoryNameRule {
	return &AwsCodecommitRepositoryInvalidRepositoryNameRule{
		resourceType:  "aws_codecommit_repository",
		attributeName: "repository_name",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w\.-]+$`),
	}
}

// Name returns the rule name
func (r *AwsCodecommitRepositoryInvalidRepositoryNameRule) Name() string {
	return "aws_codecommit_repository_invalid_repository_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodecommitRepositoryInvalidRepositoryNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodecommitRepositoryInvalidRepositoryNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodecommitRepositoryInvalidRepositoryNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodecommitRepositoryInvalidRepositoryNameRule) Check(runner tflint.Runner) error {
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
					"repository_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"repository_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w\.-]+$`),
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
