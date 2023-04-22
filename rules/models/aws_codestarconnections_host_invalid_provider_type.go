// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodestarconnectionsHostInvalidProviderTypeRule checks the pattern is valid
type AwsCodestarconnectionsHostInvalidProviderTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodestarconnectionsHostInvalidProviderTypeRule returns new rule with default attributes
func NewAwsCodestarconnectionsHostInvalidProviderTypeRule() *AwsCodestarconnectionsHostInvalidProviderTypeRule {
	return &AwsCodestarconnectionsHostInvalidProviderTypeRule{
		resourceType:  "aws_codestarconnections_host",
		attributeName: "provider_type",
		enum: []string{
			"Bitbucket",
			"GitHub",
			"GitHubEnterpriseServer",
		},
	}
}

// Name returns the rule name
func (r *AwsCodestarconnectionsHostInvalidProviderTypeRule) Name() string {
	return "aws_codestarconnections_host_invalid_provider_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodestarconnectionsHostInvalidProviderTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodestarconnectionsHostInvalidProviderTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodestarconnectionsHostInvalidProviderTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodestarconnectionsHostInvalidProviderTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as provider_type`, truncateLongMessage(val)),
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
