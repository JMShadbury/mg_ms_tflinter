PLUGIN_NAME = mg_ms_tflinter
TFLINT_CONFIG = .tflint.hcl
VERSION = 0.0.1

# Install dependencies
install:
	go mod init github.com/JMShadbury/mg_ms_tflinter
	go get github.com/terraform-linters/tflint-plugin-sdk

# Build the plugin
build:
	go build -o tflint-rules ./cmd

# Install plugin locally
install-plugin: build
	mkdir -p ~/.tflint.d/plugins
	mv tflint-rules ~/.tflint.d/plugins/tflint-rules-${PLUGIN_NAME}

# Run tests
test: install-plugin
	@echo "Testing valid configuration..."
	@tflint --config=$(TFLINT_CONFIG) testdata/valid/main.tf || { echo "Valid config test failed"; exit 1; }
	@echo "Testing invalid configuration..."
	@tflint --config=$(TFLINT_CONFIG) testdata/invalid/main.tf | grep -q "terraform_workspace_warning" || { echo "Invalid config test failed"; exit 1; }
	@echo "All tests passed!"

# Clean up
clean:
	rm -f tflint-rules

# Delete all releases from GitHub
delete-releases:
	@echo "Deleting all releases..."
	@gh release list --limit 1000 | awk '{print $$3}' | xargs -n1 gh release delete -y || true
	@echo "All releases deleted!"
