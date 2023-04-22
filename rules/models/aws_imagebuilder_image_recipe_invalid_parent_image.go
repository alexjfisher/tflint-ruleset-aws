// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsImagebuilderImageRecipeInvalidParentImageRule checks the pattern is valid
type AwsImagebuilderImageRecipeInvalidParentImageRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsImagebuilderImageRecipeInvalidParentImageRule returns new rule with default attributes
func NewAwsImagebuilderImageRecipeInvalidParentImageRule() *AwsImagebuilderImageRecipeInvalidParentImageRule {
	return &AwsImagebuilderImageRecipeInvalidParentImageRule{
		resourceType:  "aws_imagebuilder_image_recipe",
		attributeName: "parent_image",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsImagebuilderImageRecipeInvalidParentImageRule) Name() string {
	return "aws_imagebuilder_image_recipe_invalid_parent_image"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsImagebuilderImageRecipeInvalidParentImageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsImagebuilderImageRecipeInvalidParentImageRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsImagebuilderImageRecipeInvalidParentImageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsImagebuilderImageRecipeInvalidParentImageRule) Check(runner tflint.Runner) error {
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
					"parent_image must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"parent_image must be 1 characters or higher",
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
