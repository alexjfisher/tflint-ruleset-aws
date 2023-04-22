// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule checks the pattern is valid
type AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule returns new rule with default attributes
func NewAwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule() *AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule {
	return &AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule{
		resourceType:  "aws_chime_voice_connector_termination_credentials",
		attributeName: "voice_connector_id",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule) Name() string {
	return "aws_chime_voice_connector_termination_credentials_invalid_voice_connector_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsChimeVoiceConnectorTerminationCredentialsInvalidVoiceConnectorIDRule) Check(runner tflint.Runner) error {
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
