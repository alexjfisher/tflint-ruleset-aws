// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule checks the pattern is valid
type AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule returns new rule with default attributes
func NewAwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule() *AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule {
	return &AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule{
		resourceType:  "aws_apprunner_auto_scaling_configuration_version",
		attributeName: "auto_scaling_configuration_name",
		max:           32,
		min:           4,
		pattern:       regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9\-_]{3,31}$`),
	}
}

// Name returns the rule name
func (r *AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule) Name() string {
	return "aws_apprunner_auto_scaling_configuration_version_invalid_auto_scaling_configuration_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsApprunnerAutoScalingConfigurationVersionInvalidAutoScalingConfigurationNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"auto_scaling_configuration_name must be 32 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"auto_scaling_configuration_name must be 4 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[A-Za-z0-9][A-Za-z0-9\-_]{3,31}$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
