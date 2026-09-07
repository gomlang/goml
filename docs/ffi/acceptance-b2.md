# B2 acceptance audit

Scope: plan.md B2, raw Go errors, checked wrappers, explicit loss policies and bidirectional conversion. B3/B4/C/D remain outside this audit.

| Requirement | Evidence |
| --- | --- |
| Nullable ffi::Error retains Go error | Go universe error identity and builtin go_error contract; project_raw_go_errors_preserve_typed_nil_chains_and_partial_values |
| True nil differs from typed nil | error_is_nil compares the interface; raw-error and public-message regressions preserve typed-nil payload and method behavior |
| Non-null wrapper and checked conversion | NonNilError::from_error, private storage/public-interface filtering and invalid construction regressions |
| Display without replacing original error | message_bytes copies Error() bytes; message validates UTF-8; invalid bytes and independent buffer mutation tested |
| Matching preserves original errors and chains | error_matches/NonNilError::matches call native errors.Is; tests cover wrapped errors, custom Is, same text on different objects, nil and typed nil |
| Raw (T,error) remains tuple | File.Read and partial-reader fixtures bind tuples directly |
| Default Outcome preserves both values | Outcome::new and public value/error fields; actual Reader returns n > 0 and io.EOF with visible bytes |
| Result conversion is explicit with a loss policy | into_result_discarding_value_on_error names the discarded value policy; errors retain their original interface |
| Reverse conversion declares failure values | from_raw_result/from_result require caller-supplied failure_value; into_tuple returns the pair unchanged |
| Err(nil) cannot masquerade as Go success | Reverse conversion returns BoundaryError::NilError; Go-side fixture verifies the separate branch and message |
| Error chains round-trip | Raw-error package/closure transport followed by Go errors.Is; reverse-result fixture verifies sentinel identity at the Go boundary |
| Actual partial-reader acceptance | project_file_methods_close_shared_handles_and_preserve_partial_reads verifies shared bytes, partial count and EOF before and after explicit conversion |
| Packaging, navigation and stage0 | ffi depends explicitly on bytes/utf8; new query API navigation test; compiler/driver sources do not consume the new APIs; full CI required below |

Focused public message/matching test passed (/tmp/goml-error-message-test2.log). Final latest-source just ci passed with 826 compiler tests and 114 driver tests, Go helper checks, golden verification, stage3 fixed point and extracted-release smoke (/tmp/goml-error-message-final-ci.log). The public API navigation test passed. B2 is accepted against the explicit requirements above.
