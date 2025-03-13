PLUGIN_NAME = mg_ms_tflinter
TFLINT_CONFIG = .tflint.hcl

# Install dependencies
install:
	go get github.com/terraform-linters/tflint-plugin-sdk

# Run go mod tidy before building
tidy:
	go mod tidy

# Build the plugin
build: tidy
	go build -o tflint-rules ./cmd

# Run TFLint tests on all Terraform files in testdata/valid and testdata/invalid
test:
	@echo "Running tests..."
	@failed=0; \
	for file in testdata/valid/*.tf; do \
		echo "Checking valid file: $$file"; \
		if tflint --config=$(TFLINT_CONFIG) $$file 2>&1 | grep -q "Warning:"; then \
			echo "❌ Test failed: $$file should not trigger warnings."; \
			failed=1; \
		fi; \
	done; \
	for file in testdata/invalid/*.tf; do \
		echo "Checking invalid file: $$file"; \
		if ! tflint --config=$(TFLINT_CONFIG) $$file 2>&1 | grep -q "Warning:"; then \
			echo "❌ Test failed: $$file should have triggered a warning."; \
			failed=1; \
		fi; \
	done; \
	if [ $$failed -eq 0 ]; then \
		echo "✅ All tests passed."; \
	else \
		echo "❌ Some tests failed."; \
		exit 1; \
	fi

# Clean build artifacts
clean:
	rm -f tflint-rules
