# archive consumer

This independent module imports `ecosystem::archive = "0.1.0"`. Its runnable
example builds a small project bundle through the generic streaming TAR writer,
reads it back, repacks it as Deflate ZIP, checks the metadata and source payload,
and exercises the tar.gz convenience API.
Both the consumer and archive dependency use pure GoML without a `go.mod` or Go
adapter sources.

Run `just ecosystem-test archive` from the repository root to create an isolated
registry snapshot and validate the library and consumer. With that snapshot's
`GOML_HOME`, run `../../../stage2/bin/goml run` or `goml test` here.
