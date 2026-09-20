# Independent reference fixtures

Reference: [Rust bitflags 2.13.2](https://static.crates.io/crates/bitflags/bitflags-2.13.2.crate).

4,601 cases over seven profiles (access, overlap, empty, external, wide, medium, word), operations, iteration, formatting and parsing. Input generation used seed 771029. Expected lines were emitted by oracle.rs compiled with the pinned Rust crate, never by the GoML library.

Reference archive SHA-256: `3ded4057c258ba199e2d26386d3af3780957ecaee6c4ef4041c6b4b8b97c0b06`.

The fixture was captured once during migration of the verification harness. Normal tests read it directly with GoML; no Python interpreter, package download or reference runtime is required. Inputs and expected values are independent of the GoML implementation.

Fixture SHA-256: `1ecf40a7166598abdb67c81f0969853101c267a4feb3eb5c5e3678ad96217671`.
