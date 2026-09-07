# B1 acceptance audit

Scope: plan.md B1, nullable Ptr, explicit receiver method bindings, pointer identity/GC representation and os.File acceptance. This does not accept B2/B3/B4/C/D.

| Requirement | Evidence |
| --- | --- |
| Nullable ffi::Ptr[T], nil checks and controlled construction | lib/std/ffi/ffi.gom, builtin go_ptr contract and direct nil intrinsics; project_public_ffi_pointers_preserve_nil_and_go_alias_identity |
| Distinct from Ref and integers, ordinary Go pointer reachability | GoPointer through TAST/Core/Mono/repr; Go Pointer lowering; source regressions reject integer/Ref/layout construction and preserve closure/generic aliases |
| go_method first parameter is receiver | tast/env.gom attribute checks; GoMethod callable transport and artifact validation; explicit MethodExpression in Go AST |
| Actual Go value/pointer method sets | Go helper method-expression witnesses and selection metadata; TestMethodBindingsUseActualReceiverMethodSets covers value/pointer/interface/alias/promoted/generic receivers and invalid shapes |
| Methods survive packages, function values and artifacts | project_go_methods_preserve_pointer_identity_nil_and_function_values runs twice, including cached artifacts, a public child-package method and closure capture |
| os.Open, File.Read and explicit Close | project_file_methods_close_shared_handles_and_preserve_partial_reads binds the actual os functions/methods |
| Aliases refer to the same resource | Closing the alias makes Read on the original pointer return an error |
| Nil receiver retains Go behavior | nil File.Close returns its original non-nil Go error; custom IsNil method sees nil unchanged; no replacement allocation in method emission |
| No compiler-generated automatic Close or global handle table | Pointer lowering emits *T, method emission passes arguments to native Go expressions; resource fixture closes explicitly |
| Buffer policy | Existing mut_slice(...) explicitly creates the mutable view; Go writes are visible in its Vec backing storage; no automatic conversion is introduced |

Focused resource test passed in /tmp/goml-file-resource-test.log. Final latest-source just ci passed with 825 compiler tests and 113 driver tests, Go helper checks, golden verification, stage3 fixed point and extracted-release smoke (/tmp/goml-error-reverse-resource-ci.log). B1 is accepted against the explicit requirements above. B2 raw Error provides the File error representation, but B2 public message/matching adapters remain outstanding.
