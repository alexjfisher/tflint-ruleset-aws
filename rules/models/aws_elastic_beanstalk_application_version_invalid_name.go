// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElasticBeanstalkApplicationVersionInvalidNameRule checks the pattern is valid
type AwsElasticBeanstalkApplicationVersionInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkApplicationVersionInvalidNameRule returns new rule with default attributes
func NewAwsElasticBeanstalkApplicationVersionInvalidNameRule() *AwsElasticBeanstalkApplicationVersionInvalidNameRule {
	return &AwsElasticBeanstalkApplicationVersionInvalidNameRule{
		resourceType:  "aws_elastic_beanstalk_application_version",
		attributeName: "name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Name() string {
	return "aws_elastic_beanstalk_application_version_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Check(runner tflint.Runner) error {
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
					"name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
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
