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
			Version: "1.0.0",
			Rules: []tflint.Rule{
				rules.NewWorkspaceWarningRule(),
			},
		},
	})
}
