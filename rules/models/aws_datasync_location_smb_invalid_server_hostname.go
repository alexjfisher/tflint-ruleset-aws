// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsDatasyncLocationSmbInvalidServerHostnameRule checks the pattern is valid
type AwsDatasyncLocationSmbInvalidServerHostnameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsDatasyncLocationSmbInvalidServerHostnameRule returns new rule with default attributes
func NewAwsDatasyncLocationSmbInvalidServerHostnameRule() *AwsDatasyncLocationSmbInvalidServerHostnameRule {
	return &AwsDatasyncLocationSmbInvalidServerHostnameRule{
		resourceType:  "aws_datasync_location_smb",
		attributeName: "server_hostname",
		max:           255,
		pattern:       regexp.MustCompile(`^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`),
	}
}

// Name returns the rule name
func (r *AwsDatasyncLocationSmbInvalidServerHostnameRule) Name() string {
	return "aws_datasync_location_smb_invalid_server_hostname"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDatasyncLocationSmbInvalidServerHostnameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDatasyncLocationSmbInvalidServerHostnameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDatasyncLocationSmbInvalidServerHostnameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDatasyncLocationSmbInvalidServerHostnameRule) Check(runner tflint.Runner) error {
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
					"server_hostname must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`),
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
