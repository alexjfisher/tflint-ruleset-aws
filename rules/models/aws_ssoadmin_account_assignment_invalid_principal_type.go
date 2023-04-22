// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule checks the pattern is valid
type AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsoadminAccountAssignmentInvalidPrincipalTypeRule returns new rule with default attributes
func NewAwsSsoadminAccountAssignmentInvalidPrincipalTypeRule() *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule {
	return &AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule{
		resourceType:  "aws_ssoadmin_account_assignment",
		attributeName: "principal_type",
		enum: []string{
			"USER",
			"GROUP",
		},
	}
}

// Name returns the rule name
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Name() string {
	return "aws_ssoadmin_account_assignment_invalid_principal_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsoadminAccountAssignmentInvalidPrincipalTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as principal_type`, truncateLongMessage(val)),
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
