// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSagemakerDomainInvalidAppNetworkAccessTypeRule checks the pattern is valid
type AwsSagemakerDomainInvalidAppNetworkAccessTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSagemakerDomainInvalidAppNetworkAccessTypeRule returns new rule with default attributes
func NewAwsSagemakerDomainInvalidAppNetworkAccessTypeRule() *AwsSagemakerDomainInvalidAppNetworkAccessTypeRule {
	return &AwsSagemakerDomainInvalidAppNetworkAccessTypeRule{
		resourceType:  "aws_sagemaker_domain",
		attributeName: "app_network_access_type",
		enum: []string{
			"PublicInternetOnly",
			"VpcOnly",
		},
	}
}

// Name returns the rule name
func (r *AwsSagemakerDomainInvalidAppNetworkAccessTypeRule) Name() string {
	return "aws_sagemaker_domain_invalid_app_network_access_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSagemakerDomainInvalidAppNetworkAccessTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSagemakerDomainInvalidAppNetworkAccessTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSagemakerDomainInvalidAppNetworkAccessTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSagemakerDomainInvalidAppNetworkAccessTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as app_network_access_type`, truncateLongMessage(val)),
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
