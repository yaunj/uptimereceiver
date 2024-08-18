# uptimereceiver

Simple metrics receiver to gather system uptime as a metric.


## Developing

If you have `make` installed, you can test the code like this:

```bash
make generate
make test
```

Generate parts of the receiver using the mdatagen tool:

```bash
cd uptimereceiver
go1.21.0 install go.opentelemetry.io/collector/cmd/mdatagen
go1.21.0 generate ./...
```

Download [ocb](https://opentelemetry.io/docs/collector/custom-collector/) and
create the test collector:

```bash
curl --proto '=https' --tlsv1.2 -fL -o ocb https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd%2Fbuilder%2Fv0.107.0/ocb_0.107.0_linux_amd64
chmod +x ocb
./ocb --config builder-config.yaml
```

You can now run the collector like this: `./otelcol-dev/otelcol-dev --config
config.yaml`


## Resources

* [Building a custom
  collector](https://opentelemetry.io/docs/collector/custom-collector/)
* [Building a
  receiver](https://opentelemetry.io/docs/collector/building/receiver/) -
  focuses on a receiver for traces
  * I had issues using the Go workspace and manual changes to the generated
    collector code, as suggested in the tutorial. In stead I added my receiver
    to the `builder-config.yaml` with a local path. Documentation for the
    builder config can be found
    [here](https://pkg.go.dev/go.opentelemetry.io/collector/cmd/builder#readme-configuration)
* [OpenTelemetry: Building a Custom
  Collector](https://mmynk.com/tech/opentelemetry.html) looks interesting for
  more details on `mdatagen` and metrics specifics.
