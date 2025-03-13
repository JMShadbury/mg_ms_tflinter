package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/hashicorp/hcl/v2"
)

// WorkspaceWarningRule detects the use of terraform.workspace
type WorkspaceWarningRule struct {
	tflint.DefaultRule
}

// NewWorkspaceWarningRule initializes the rule
func NewWorkspaceWarningRule() *WorkspaceWarningRule {
	return &WorkspaceWarningRule{}
}

// Name returns the rule name
func (r *WorkspaceWarningRule) Name() string {
	return "terraform_workspace_warning"
}

// Enabled ensures the rule is enabled by default
func (r *WorkspaceWarningRule) Enabled() bool {
	return true
}

// Severity returns the rule severity (Warning)
func (r *WorkspaceWarningRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link provides documentation for this rule
func (r *WorkspaceWarningRule) Link() string {
	return "https://developer.hashicorp.com/terraform/language/state/workspaces"
}

// Check runs the rule logic
func (r *WorkspaceWarningRule) Check(runner tflint.Runner) error {
	expressions, err := runner.GetModuleExpressions()
	if err != nil {
		return err
	}

	for _, expr := range expressions {
		// Check if expression is "terraform.workspace"
		if expr.Traversal.RootName() == "terraform" &&
			len(expr.Traversal) > 1 &&
			expr.Traversal[1].(hcl.TraverseAttr).Name == "workspace" {
			runner.EmitIssue(r, "Warning: Avoid using terraform.workspace as it can lead to unexpected behavior.", expr.Range)
		}
	}

	return nil
}
