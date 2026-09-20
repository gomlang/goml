# bench consumer

An independent `ecosystem::bench = "0.1.0"` consumer that compares insertion sort
and standard sorting at 32 and 128 elements. Inputs pass through the scalar
black-box boundary, setup allocates fresh vectors outside measured work, and
returned digests retain each algorithm's result. The example checks algorithmic
agreement without imposing performance thresholds.

`goml run` writes raw samples/configuration and summaries to
`_artifact/bench-report.json` and a standalone HTML report to
`_artifact/bench-report.html`, then reloads the JSON and compares a persisted
baseline. The consumer test uses an injected clock to exercise cross-module
setup closures and statistical/reporting APIs deterministically.

Run `just ecosystem-test bench` from the repository root to construct the
isolated registry. Use the reported `GOML_HOME` for direct `goml check`, `goml test`
or `goml run` commands here. The Go manifest resolves only the small local atomic
adapter; benchmark/statistical source is resolved through the versioned GoML
registry.
