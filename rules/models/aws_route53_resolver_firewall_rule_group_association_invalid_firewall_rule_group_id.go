// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule checks the pattern is valid
type AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule returns new rule with default attributes
func NewAwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule() *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule {
	return &AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule{
		resourceType:  "aws_route53_resolver_firewall_rule_group_association",
		attributeName: "firewall_rule_group_id",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule) Name() string {
	return "aws_route53_resolver_firewall_rule_group_association_invalid_firewall_rule_group_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverFirewallRuleGroupAssociationInvalidFirewallRuleGroupIDRule) Check(runner tflint.Runner) error {
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
					"firewall_rule_group_id must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"firewall_rule_group_id must be 1 characters or higher",
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
