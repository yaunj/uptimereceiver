# uptimereceiver

Simple metrics receiver to gather system uptime as a metric.


## Developing

Generate parts of the receiver using the mdatagen tool:

```bash
cd uptimereceiver
go1.21.0 install go.opentelemetry.io/collector/cmd/mdatagen
go1.21.0 generate ./...
```

Download [https://opentelemetry.io/docs/collector/custom-collector/](ocb) and
create the test collector:

```bash
curl --proto '=https' --tlsv1.2 -fL -o ocb https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd%2Fbuilder%2Fv0.107.0/ocb_0.107.0_linux_amd64
chmod +x ocb
./ocb --config builder-config.yaml
```

You can now run the collector like this: `./otelcol-dev/otelcol-dev --config
config.yaml`
