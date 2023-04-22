// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAPIGatewayGatewayResponseInvalidResponseTypeRule checks the pattern is valid
type AwsAPIGatewayGatewayResponseInvalidResponseTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayGatewayResponseInvalidResponseTypeRule returns new rule with default attributes
func NewAwsAPIGatewayGatewayResponseInvalidResponseTypeRule() *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule {
	return &AwsAPIGatewayGatewayResponseInvalidResponseTypeRule{
		resourceType:  "aws_api_gateway_gateway_response",
		attributeName: "response_type",
		enum: []string{
			"DEFAULT_4XX",
			"DEFAULT_5XX",
			"RESOURCE_NOT_FOUND",
			"UNAUTHORIZED",
			"INVALID_API_KEY",
			"ACCESS_DENIED",
			"AUTHORIZER_FAILURE",
			"AUTHORIZER_CONFIGURATION_ERROR",
			"INVALID_SIGNATURE",
			"EXPIRED_TOKEN",
			"MISSING_AUTHENTICATION_TOKEN",
			"INTEGRATION_FAILURE",
			"INTEGRATION_TIMEOUT",
			"API_CONFIGURATION_ERROR",
			"UNSUPPORTED_MEDIA_TYPE",
			"BAD_REQUEST_PARAMETERS",
			"BAD_REQUEST_BODY",
			"REQUEST_TOO_LARGE",
			"THROTTLED",
			"QUOTA_EXCEEDED",
			"WAF_FILTERED",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Name() string {
	return "aws_api_gateway_gateway_response_invalid_response_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayGatewayResponseInvalidResponseTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as response_type`, truncateLongMessage(val)),
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
