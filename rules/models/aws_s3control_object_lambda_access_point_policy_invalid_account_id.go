// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule checks the pattern is valid
type AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule returns new rule with default attributes
func NewAwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule() *AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule {
	return &AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule{
		resourceType:  "aws_s3control_object_lambda_access_point_policy",
		attributeName: "account_id",
		max:           64,
		pattern:       regexp.MustCompile(`^\d{12}$`),
	}
}

// Name returns the rule name
func (r *AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule) Name() string {
	return "aws_s3control_object_lambda_access_point_policy_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3controlObjectLambdaAccessPointPolicyInvalidAccountIDRule) Check(runner tflint.Runner) error {
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
					"account_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\d{12}$`),
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
