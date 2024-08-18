GO?=go1.21.0

ocb:
	@echo Installing ocb
	@curl -s --proto '=https' --tlsv1.2 -fL -o ocb \
	https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd%2Fbuilder%2Fv0.107.0/ocb_0.107.0_linux_amd64
	@chmod +x ocb

otelcol-dev/otelcol-dev: ocb builder-config.yaml
	./ocb --config builder-config.yaml

.PHONY: mdatagen
mdatagen:
	cd uptimereceiver && $(GO) install go.opentelemetry.io/collector/cmd/mdatagen

.PHONY: generate
generate: mdatagen
	cd uptimereceiver && $(GO) generate ./...

.PHONY: test
test: otelcol-dev/otelcol-dev
	./otelcol-dev/otelcol-dev --config config.yaml
