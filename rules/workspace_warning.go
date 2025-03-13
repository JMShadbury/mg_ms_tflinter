package rules

import (
    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// WorkspaceWarningRule checks for terraform.workspace usage
type WorkspaceWarningRule struct {
    tflint.DefaultRule
}

// NewWorkspaceWarningRule returns a new rule
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

// Severity returns the rule severity
func (r *WorkspaceWarningRule) Severity() tflint.Severity {
    return tflint.WARNING
}

// Link returns documentation link
func (r *WorkspaceWarningRule) Link() string {
    return "https://developer.hashicorp.com/terraform/language/state/workspaces"
}

// Check inspects Terraform configuration for terraform.workspace usage
func (r *WorkspaceWarningRule) Check(runner tflint.Runner) error {
    // Get all module content
    body, err := runner.GetModuleContent(&hclext.BodySchema{
        Attributes: []hclext.AttributeSchema{{Name: "*"}},
    }, &tflint.GetModuleContentOption{})
    if err != nil {
        return err
    }

    // Walk through all expressions
    for _, attr := range body.Attributes {
        err := hclext.WalkExpressions(attr.Expr, func(expr hclext.Expression) (bool, error) {
            // Check if the expression contains terraform.workspace
            vars := expr.Variables()
            for _, v := range vars {
                if len(v) > 1 && v[0].String() == "terraform" && v[1].String() == "workspace" {
                    err := runner.EmitIssue(
                        r,
                        "Warning: Avoid using terraform.workspace as it can lead to unexpected behavior",
                        v.Range(),
                    )
                    if err != nil {
                        return false, err
                    }
                }
            }
            return true, nil
        })
        if err != nil {
            return err
        }
    }

    return nil
}