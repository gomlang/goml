# Config consumer

This executable resolves `ecosystem::config = "0.1.0"` as a normal registry
dependency. It combines typed defaults, TOML, environment entries and explicit
CLI overrides, validates typed settings, queries field provenance, and checks
that saved snapshots survive a live reload.

Run `just ecosystem-test config` from the repository root once the module is
registered with the native verifier. The consumer also runs through `goml test`
and `goml run` using the verifier's isolated registry home.
