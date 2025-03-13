PLUGIN_NAME = mg_ms_tflinter
TFLINT_CONFIG = .tflint.hcl
VERSION = 0.0.2

# Install dependencies
install:
	go mod init github.com/JMShadbury/mg_ms_tflinter || true
	go get github.com/terraform-linters/tflint-plugin-sdk
	go get github.com/hashicorp/hcl/v2
	go mod tidy
	go build -o tflint-rules ./cmd

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
	go clean -modcache
	rm -f tflint-rules
	rm -f go.sum
	rm -f go.mod

release: install build
	git tag -a v$(VERSION) -m "Release v$(VERSION)"
	git push origin v$(VERSION)
	gh release create v$(VERSION) --title "Release v$(VERSION)" --notes "Release notes for v$(VERSION)" --draft
	@echo "Release v$(VERSION) created!"

# Delete all releases from GitHub
delete-releases:
	@echo "Deleting all releases and tags..."
	@gh release list --limit 1000 | awk '{print $$3}' | xargs -I {} sh -c 'gh release delete -y {}; git push origin :refs/tags/{}; git tag -d {} || true' || true
	@echo "All releases and tags deleted!"
