// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule checks the pattern is valid
type AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule returns new rule with default attributes
func NewAwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule() *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule {
	return &AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule{
		resourceType:  "aws_worklink_website_certificate_authority_association",
		attributeName: "display_name",
		max:           100,
	}
}

// Name returns the rule name
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule) Name() string {
	return "aws_worklink_website_certificate_authority_association_invalid_display_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWorklinkWebsiteCertificateAuthorityAssociationInvalidDisplayNameRule) Check(runner tflint.Runner) error {
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
					"display_name must be 100 characters or less",
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
