// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCodepipelineInvalidRoleArnRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_codepipeline" "foo" {
	role_arn = "arn:aws:iam::123456789012:instance-profile/s3access-profile"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCodepipelineInvalidRoleArnRule(),
					Message: `role_arn does not match valid pattern ^arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_codepipeline" "foo" {
	role_arn = "arn:aws:iam::123456789012:role/s3access"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCodepipelineInvalidRoleArnRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
