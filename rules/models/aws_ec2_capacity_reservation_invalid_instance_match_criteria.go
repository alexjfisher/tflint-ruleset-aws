// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule checks the pattern is valid
type AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule returns new rule with default attributes
func NewAwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule() *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule {
	return &AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule{
		resourceType:  "aws_ec2_capacity_reservation",
		attributeName: "instance_match_criteria",
		enum: []string{
			"open",
			"targeted",
		},
	}
}

// Name returns the rule name
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Name() string {
	return "aws_ec2_capacity_reservation_invalid_instance_match_criteria"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEc2CapacityReservationInvalidInstanceMatchCriteriaRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as instance_match_criteria`, truncateLongMessage(val)),
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
