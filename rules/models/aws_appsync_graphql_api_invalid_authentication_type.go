// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule checks the pattern is valid
type AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule returns new rule with default attributes
func NewAwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule() *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule {
	return &AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule{
		resourceType:  "aws_appsync_graphql_api",
		attributeName: "authentication_type",
		enum: []string{
			"API_KEY",
			"AWS_IAM",
			"AMAZON_COGNITO_USER_POOLS",
			"OPENID_CONNECT",
			"AWS_LAMBDA",
		},
	}
}

// Name returns the rule name
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Name() string {
	return "aws_appsync_graphql_api_invalid_authentication_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as authentication_type`, truncateLongMessage(val)),
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
