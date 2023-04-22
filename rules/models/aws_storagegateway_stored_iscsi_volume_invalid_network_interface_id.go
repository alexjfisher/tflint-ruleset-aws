// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule checks the pattern is valid
type AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule returns new rule with default attributes
func NewAwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule() *AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule {
	return &AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule{
		resourceType:  "aws_storagegateway_stored_iscsi_volume",
		attributeName: "network_interface_id",
		pattern:       regexp.MustCompile(`^\A(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}\z$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule) Name() string {
	return "aws_storagegateway_stored_iscsi_volume_invalid_network_interface_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayStoredIscsiVolumeInvalidNetworkInterfaceIDRule) Check(runner tflint.Runner) error {
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
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^\A(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}\z$`),
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
