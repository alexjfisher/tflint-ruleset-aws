// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsNetworkfirewallResourcePolicyInvalidPolicyRule checks the pattern is valid
type AwsNetworkfirewallResourcePolicyInvalidPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsNetworkfirewallResourcePolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsNetworkfirewallResourcePolicyInvalidPolicyRule() *AwsNetworkfirewallResourcePolicyInvalidPolicyRule {
	return &AwsNetworkfirewallResourcePolicyInvalidPolicyRule{
		resourceType:  "aws_networkfirewall_resource_policy",
		attributeName: "policy",
		max:           395000,
		min:           1,
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Name() string {
	return "aws_networkfirewall_resource_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkfirewallResourcePolicyInvalidPolicyRule) Check(runner tflint.Runner) error {
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
					"policy must be 395000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"policy must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*\S.*$`),
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
