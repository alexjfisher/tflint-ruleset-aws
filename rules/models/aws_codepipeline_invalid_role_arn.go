// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodepipelineInvalidRoleArnRule checks the pattern is valid
type AwsCodepipelineInvalidRoleArnRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsCodepipelineInvalidRoleArnRule returns new rule with default attributes
func NewAwsCodepipelineInvalidRoleArnRule() *AwsCodepipelineInvalidRoleArnRule {
	return &AwsCodepipelineInvalidRoleArnRule{
		resourceType:  "aws_codepipeline",
		attributeName: "role_arn",
		max:           1024,
		pattern:       regexp.MustCompile(`^arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*$`),
	}
}

// Name returns the rule name
func (r *AwsCodepipelineInvalidRoleArnRule) Name() string {
	return "aws_codepipeline_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodepipelineInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodepipelineInvalidRoleArnRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodepipelineInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodepipelineInvalidRoleArnRule) Check(runner tflint.Runner) error {
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
					"role_arn must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*$`),
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
