package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// WorkspaceWarningRule checks the use of terraform.workspace
type WorkspaceWarningRule struct {
	tflint.DefaultRule
}

// NewWorkspaceWarningRule returns a new WorkspaceWarningRule
func NewWorkspaceWarningRule() *WorkspaceWarningRule {
	return &WorkspaceWarningRule{}
}

// Name returns the rule name
func (r *WorkspaceWarningRule) Name() string {
	return "terraform_workspace_warning"
}

// Enabled returns whether the rule is enabled by default
func (r *WorkspaceWarningRule) Enabled() bool {
	return true
}

// Severity returns the rule severity (Warning)
func (r *WorkspaceWarningRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link provides a link to documentation for this rule
func (r *WorkspaceWarningRule) Link() string {
	return "https://developer.hashicorp.com/terraform/language/state/workspaces"
}

// Check checks whether terraform.workspace is used
func (r *WorkspaceWarningRule) Check(runner tflint.Runner) error {
	// This rule is an example to get attributes of blocks.
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "terraform",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type: "workspace",
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, terraform := range content.Blocks {
		for _, workspace := range terraform.Body.Blocks {
			err := runner.EmitIssue(
				r,
				"Warning: Avoid using terraform.workspace as it can lead to unexpected behavior.",
				workspace.DefRange,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
