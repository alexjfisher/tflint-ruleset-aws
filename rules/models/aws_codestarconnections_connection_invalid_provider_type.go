// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodestarconnectionsConnectionInvalidProviderTypeRule checks the pattern is valid
type AwsCodestarconnectionsConnectionInvalidProviderTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCodestarconnectionsConnectionInvalidProviderTypeRule returns new rule with default attributes
func NewAwsCodestarconnectionsConnectionInvalidProviderTypeRule() *AwsCodestarconnectionsConnectionInvalidProviderTypeRule {
	return &AwsCodestarconnectionsConnectionInvalidProviderTypeRule{
		resourceType:  "aws_codestarconnections_connection",
		attributeName: "provider_type",
		enum: []string{
			"Bitbucket",
			"GitHub",
			"GitHubEnterpriseServer",
			"GitLab",
			"GitLabSelfManaged",
		},
	}
}

// Name returns the rule name
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Name() string {
	return "aws_codestarconnections_connection_invalid_provider_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodestarconnectionsConnectionInvalidProviderTypeRule) Check(runner tflint.Runner) error {
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
