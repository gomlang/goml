# Compile-time evaluation architecture

GoML compile-time evaluation is a typed frontend phase:

```text
lexer → parser → CST → AST → HIR → TAST
                                      │
                                      ├─ capability validation
                                      ├─ CTIR lowering
dependency interfaces and CTIR ───────┤
                                      ├─ CTIR verification
                                      ├─ deterministic evaluation
                                      └─ value reification
                                                │
                                                ▼
                                      finalized TAST
                                                │
                           Core → Mono → Lift → ANF → Go
```

The phase boundary has three invariants:

- `#[comptime]` bodies are validated in their defining package even when no call site currently uses them.
- Every imported CTIR unit is treated as untrusted data and verified before it enters a dependency environment.
- Finalized TAST contains no `Comptime` expression. Core lowering checks this invariant again, so Core and every later IR are independent of CTFE.

CTIR is a typed, direct-call, tree-walking representation. It contains deterministic values, local state, structured control flow, pattern matching, fixed-array and builtin-range iteration, direct calls, deterministic string and integer-formatting intrinsics, and `compile_error`. It has no closures, indirect calls, dynamic dispatch, runtime hooks, host I/O, references, channels, or concurrency. Local slot types, call signatures, control-flow targets, source-origin IDs, structured value types, and target compatibility are verified before evaluation.

Evaluation uses semantic fuel, call-depth, temporary-node, temporary-memory, and final-result-size limits. It never uses a wall-clock deadline. Direct-call results are memoized only within the current evaluation. Integer values store signedness, width, and raw bits; the CTIR target specification determines the width of `isize` and `usize` and participates in semantic hashing.

Integer `to_string` lowers to a typed `comptime.to_string.iN` or `comptime.to_string.uN` intrinsic for widths 8, 16, 32, or 64. The verifier checks the argument width and signedness; evaluation formats raw bits using target integer semantics, including signed minima and unsigned maxima. This does not enable arbitrary trait calls in CTFE.

An interface artifact exports public comptime entries and the reachable closure of private comptime helpers and constants. Public constants carry canonical values. Source origins are debug metadata, use package-relative paths, and do not participate in the CTIR semantic hash. Function bodies, local slot types, target semantics, public values, and referenced dependency interface hashes do participate. The artifact decoder checks its format and semantic hash, and dependency loading then runs the CTIR verifier with the complete imported module set.

CTFE failures are recoverable compiler diagnostics. They identify the failing source origin, the requesting comptime site, and the compile-time call stack. A failed or resource-exhausted evaluation never terminates the compiler process.

## Programmable derive phase

Programmable derive reuses verified CTIR but runs at an earlier consumer-side phase:

```text
dependency interface → verify derive CTIR ──────────────┐
                                                        │
source → parser → CST → AST → resolve derive entry → evaluate handler
                                                        │
                                                        ▼
                                                generated AST impl
                                                        │
                                                        ▼
                                          HIR → TAST → ordinary CTFE
```

The handler was type checked and lowered to CTIR when its defining package was compiled. The consuming package never executes untyped source or host code. Evaluation receives an opaque `DeriveInput` handle and must return an opaque `DeriveOutput` handle. Compiler intrinsics expose structured attributes and type shapes through opaque `MetaAttribute` and `MetaType` handles, and expression, pattern, arm, block, and method builders construct a structured implementation. The result cannot contain arbitrary tokens or declarations, and it enters normal HIR lowering, name resolution, coherence checking, type checking, monomorphization, and code generation.

Public `#[comptime_derive]` entries and the reachable closure of private derive helpers are part of the interface CTIR semantic section. A derive body change therefore changes the interface hash. Derive entries are not runtime exports. The interface decoder verifies meta types, intrinsic signatures, direct-call targets, IDs, and the semantic hash before evaluation.

Definition-site builders qualify unqualified trait, type, and function names with the handler package. Explicit call-site builders leave names in the target package scope. Generated local bindings use handler-selected names; `derive_fresh_name` provides collision-free compiler names. Generated nodes use the requesting derive attribute as their diagnostic origin.

Derive evaluation uses the normal fuel, depth, value-node, and temporary-memory limits. It additionally permits at most 100,000 metadata or syntax-builder operations. Derive calls are not memoized because their opaque arena handles are evaluation-local.

The query layer retains formatted post-expansion AST per source file. `gomlc run-single --dump-expanded-ast` and the LSP `goml/expandedDerive` request expose the same expansion boundary without bypassing normal derive evaluation or diagnostics.

A derive output contains one trait implementation or one inherent implementation. `derive_output_inherent(input)` preserves the target's generic parameters and identity. `derive_output_add_method` keeps inherent methods private, while `derive_output_add_public_method` marks an inherent method public and diagnoses use on a trait implementation. Both enter the usual visibility and type-checking pipeline. Associated constants, multiple output impls, and arbitrary top-level declarations are still outside the output contract.
