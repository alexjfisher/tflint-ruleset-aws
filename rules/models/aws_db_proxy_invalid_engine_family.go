// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDBProxyInvalidEngineFamilyRule checks the pattern is valid
type AwsDBProxyInvalidEngineFamilyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDBProxyInvalidEngineFamilyRule returns new rule with default attributes
func NewAwsDBProxyInvalidEngineFamilyRule() *AwsDBProxyInvalidEngineFamilyRule {
	return &AwsDBProxyInvalidEngineFamilyRule{
		resourceType:  "aws_db_proxy",
		attributeName: "engine_family",
		enum: []string{
			"MYSQL",
			"POSTGRESQL",
			"SQLSERVER",
		},
	}
}

// Name returns the rule name
func (r *AwsDBProxyInvalidEngineFamilyRule) Name() string {
	return "aws_db_proxy_invalid_engine_family"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDBProxyInvalidEngineFamilyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDBProxyInvalidEngineFamilyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDBProxyInvalidEngineFamilyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDBProxyInvalidEngineFamilyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as engine_family`, truncateLongMessage(val)),
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
