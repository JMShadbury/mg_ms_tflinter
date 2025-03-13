package main

import (
    "github.com/terraform-linters/tflint-plugin-sdk/plugin"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/JMShadbury/mg_ms_tflinter/rules"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        RuleSet: &tflint.BuiltinRuleSet{
            Name:    "mg_ms_tflinter",
            Version: "0.0.3", // Updated to match Makefile
            Rules: []tflint.Rule{
                rules.NewWorkspaceWarningRule(),
            },
        },
    })
}