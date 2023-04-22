// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53RecordInvalidTypeRule checks the pattern is valid
type AwsRoute53RecordInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsRoute53RecordInvalidTypeRule returns new rule with default attributes
func NewAwsRoute53RecordInvalidTypeRule() *AwsRoute53RecordInvalidTypeRule {
	return &AwsRoute53RecordInvalidTypeRule{
		resourceType:  "aws_route53_record",
		attributeName: "type",
		enum: []string{
			"SOA",
			"A",
			"TXT",
			"NS",
			"CNAME",
			"MX",
			"NAPTR",
			"PTR",
			"SRV",
			"SPF",
			"AAAA",
			"CAA",
			"DS",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53RecordInvalidTypeRule) Name() string {
	return "aws_route53_record_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53RecordInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53RecordInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53RecordInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53RecordInvalidTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
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
