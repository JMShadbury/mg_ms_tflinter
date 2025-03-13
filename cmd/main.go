package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/JMShadbury/mg_ms_tflinter/rules"
)

func main() {
	tflint.Main(tflint.Plugin{
		Rules: []tflint.Rule{
			rules.NewWorkspaceWarningRule(),
		},
	})
}
