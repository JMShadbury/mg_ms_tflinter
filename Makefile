PLUGIN_NAME = mg_ms_tflinter
TFLINT_CONFIG = .tflint.hcl
VERSION = 0.0.4

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
	@TFLINT_CONFIG=$(PWD)/$(TFLINT_CONFIG) tflint --chdir=$(PWD)/testdata/valid --filter="*" || { echo "Valid config test failed"; exit 1; }
	@echo "Testing invalid configuration..."
	@TFLINT_CONFIG=$(PWD)/$(TFLINT_CONFIG) tflint --chdir=$(PWD)/testdata/invalid --filter="*" || { echo "Invalid config test failed"; exit 1; }
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

	GOOS=darwin GOARCH=amd64 go build -o tflint-ruleset-mg_ms_tflinter_darwin_amd64
	GOOS=linux GOARCH=amd64 go build -o tflint-ruleset-mg_ms_tflinter_linux_amd64
	GOOS=windows GOARCH=amd64 go build -o tflint-ruleset-mg_ms_tflinter_windows_amd64.exe

	shasum -a 256 tflint-ruleset-mg_ms_tflinter_darwin_amd64 tflint-ruleset-mg_ms_tflinter_linux_amd64 tflint-ruleset-mg_ms_tflinter_windows_amd64.exe > checksums.txt

	gh release create v$(VERSION) --title "Release v$(VERSION)" --notes "Release notes for v$(VERSION)" --draft \
		tflint-ruleset-mg_ms_tflinter_darwin_amd64 tflint-ruleset-mg_ms_tflinter_linux_amd64 tflint-ruleset-mg_ms_tflinter_windows_amd64.exe checksums.txt

	@echo "Release v$(VERSION) created!"



# Delete all releases from GitHub
delete-releases:
	@echo "Deleting all releases and tags..."
	@gh release list --limit 1000 | awk '{print $$3}' | xargs -I {} sh -c 'gh release delete -y {}; git push origin :refs/tags/{}; git tag -d {} || true' || true
	@echo "All releases and tags deleted!"
