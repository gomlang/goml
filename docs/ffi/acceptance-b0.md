# B0 external type identity acceptance

Scope is the explicit B0 contract in `plan.md`: external named type identity throughout the compiler, recoverable rejection of unsupported shapes, nominality, cross-package identity, preserved Go method sets, artifact round trips and bootstrap. B1 pointer/method syntax and later batches are separate work.

| Requirement | Evidence |
| --- | --- |
| `#[go_type("pkg", "Name")] pub extern type Name;` | `parser::parses_external_type_declarations_and_generic_parameters`, `parser::recovers_invalid_external_type_bodies_and_missing_names`, `formatter::formats_external_type_declarations`; source module tests in `goml/cmd/goml/ffi_test.gom` |
| Canonical import path, object name and instantiated arguments | `Type::External` in TAST/Core/Mono/repr; `tast::external_types_preserve_nominality_and_inference_occurs_checks`; the external pipeline and generic artifact tests |
| `time.Duration` differs from `i64` | The nominality test explicitly checks both equality and unification; source-checking tests reject returning an external Duration as i64 or constructing it from zero |
| Same Go type imported under different GoML names/packages is equal | `tast::external_type_environment_preserves_source_identity_and_checks_arity`, `external_source_checking_preserves_identity_across_files_and_rejects_construction`, and driver `project_external_types_build_across_packages_and_revalidate_cached_aliases` |
| Alias, defined type and interface identities remain distinct where appropriate | `ffi::go_identity_normalizes_aliases_without_erasing_named_types`; live Go metadata tests load Duration, io.Reader, aliases and generic declarations in one Go package-loading world |
| No layout construction or private field access | Source-checking regression rejects `value.hidden` and `Delay {}`; external types are not GoML struct definitions |
| Original Go named type and method set survive emission | `external_type_identity_survives_artifacts_specialization_and_go_emission` round-trips Core, runs Mono/Lift/ANF/Go lowering and compiles a direct `unique.Handle[int64].Value()` call against the emitted type |
| Artifact identity and arguments survive | `external_type_interfaces_roundtrip_arguments_and_reject_missing_identity`; Core/interface metadata tests; source-free declaration and concrete-instance revalidation in required/off modes |
| Unsupported shapes fail recoverably | Source/environment tests cover bad arity, absent metadata, forbidden construction and nominal mismatch; graph/identity decoders reject unsupported raw forms and malformed metadata |
| New and old compilation stages | Repository stage0 builds and full CI stage3 fixed-point comparison; final current-state validation is recorded below |

Concrete generic source applications, original Go alias constraints, symbolic Core application transport and linked/cached specialization checks are implemented in addition to the basic declaration contract. Source applications such as `Box[T]` inside a generic GoML function still produce the documented specialization-required diagnostic. Query/LSP Go metadata scheduling and inaccessible Go alias spelling also remain limitations. These are not being described as supported by this acceptance report.

The B0 text explicitly allows recoverable rejection of unsupported shapes. General symbolic source applications are therefore not treated as a prerequisite to starting the explicitly requested B1 pointer and method work. The complete `plan.md` goal remains unfinished.

## Current validation

`just ci` completed successfully for the current implementation: 814 compiler tests, 107 driver tests, Go helper checks, stage3 fixed-point comparison and extracted-release smoke. The run includes the source-free linked specialization regression and fresh/cached project link scheduling regression. Log: `/tmp/goml-cached-specialization-ci.log`. `gomlc` and `goml` formatting and `git diff --check` passed. No stage0 advancement or release is claimed.
