package rules

import (
    "github.com/terraform-linters/tflint-plugin-sdk/hclext"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/hashicorp/hcl/v2"
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

    // Check each attribute for terraform.workspace references
    for _, attr := range body.Attributes {
        vars := attr.Expr.Variables()
        for _, traversal := range vars {
            if len(traversal) >= 2 {
                root, ok1 := traversal[0].(hcl.TraverseRoot)
                attrName, ok2 := traversal[1].(hcl.TraverseAttr)
                if ok1 && ok2 && root.Name == "terraform" && attrName.Name == "workspace" {
                    err := runner.EmitIssue(
                        r,
                        "Warning: Avoid using terraform.workspace as it can lead to unexpected behavior",
                        attr.Range, // Use attribute range instead of traversal range
                    )
                    if err != nil {
                        return err
                    }
                }
            }
        }
    }

    return nil
}