# Batch B external type contract

This is the implementation contract for plan B0, following the completed A acceptance gate. Declaration syntax, metadata-driven TAST checking, concrete generic instances, artifact transport and native Go emission are implemented. The requirement-by-requirement acceptance evidence is in [acceptance-b0.md](acceptance-b0.md); later boundaries and adapters are tracked in their respective acceptance records.

## Declaration and identity

`#[go_type("import/path", "ExportedName")] pub extern type LocalName;` declares a local spelling for an actual Go type. Visibility follows existing top-level visibility. It cannot have a GoML body, fields, variants or a layout-based constructor. Invalid attribute placement, missing/duplicate attributes, invalid exported Go object names and references to non-type objects produce recoverable diagnostics.

The type's semantic identity is its canonical Go import path, declaration object name and instantiated type arguments. Two GoML declarations referring to the same Go type in one build world unify even if their GoML package names differ. Selected module versions, replace paths, cache paths and helper-session node IDs are build provenance, not nominal identity.

Go aliases normalize to their targets. A defined Go type remains nominal even when its underlying type is a supported scalar: time.Duration does not unify with i64. An alias of Duration unifies with Duration; an alias of a primitive normalizes to that primitive's raw bridge representation. Defined interfaces remain external Go interfaces, not GoML dyn trait objects. Their identity and method sets must survive artifact roundtrips.

Generic external declarations use the language's square-bracket type parameter syntax. Applications supply Go type arguments and are checked against the actual Go declaration's arity and constraints; an uninstantiated generic Go type is not a concrete value type. Identity and substitution preserve instantiated arguments. Source applications must currently have concrete arguments; a symbolic application such as `Box[T]` inside a generic GoML function or alias reports that specialization is required.

## Metadata and checking order

A's post-typecheck call witnesses are insufficient to establish external alias identity. The orchestration layer must collect external type requests, resolve them in the selected Go build/caller context and supply structured metadata before type-directed equality and inference use those types. TAST and environment operations stay pure; they consume resolved descriptions and report missing/unsupported metadata instead of loading host packages.

Extend the helper's bounded protocol with explicit type requests and results. Reuse the same packages.Load world and structured referenced-type graph for declarations and call witnesses; do not compare go/types values across independent loads. Type-only declarations must be validated even when no function binding uses them. Cache hits require current-world validation, as in A.

Keep declared spelling/provenance separate from normalized semantic identity. Local source locations remain available for diagnostics; portable semantic hashes exclude host paths. The artifact format records enough resolved kind/identity/argument information for checking and linking without the original source. The helper's display signature is diagnostic text, never a type key.

## Compiler and backend obligations

Parser/CST/AST represent extern type as a declaration distinct from extern fn and transparent type aliases. HIR resolves its local name and visibility. TAST tracks external identity explicitly, including equality, substitution, printing and unsupported-operation diagnostics. Core, Mono, Lift, ANF/repr and their codecs preserve the identity and arguments through monomorphization and linking.

The Go backend emits the original Go type selector or a Go alias (`type Local = pkg.Type`), preserving method sets. It must not create a new Go defined type for a binding. External type-only imports participate in import naming, collision checks and DCE. GoML cannot construct an external struct by copying its layout or access private Go fields. Pointer/method operations belong to B1 and must use the same identity infrastructure.

## Required regression sequence

1. Parser, formatter and attachment diagnostics for the declaration and generic parameter form.
2. Structured helper queries for Duration, an alias of Duration, a primitive alias, an exported interface and a generic type application; missing/private/non-type/invalid-argument failures.
3. TAST rejects Duration as i64 and unifies two package-local bindings to Duration. Imported aliases preserve actual Go alias semantics.
4. Artifact roundtrip and source/artifact equivalence for nominal identity, arguments, kind and method sets; invalid artifact data fails recoverably.
5. Actual Go calls return and accept an external value through GoML, with generated Go using its original nominal type. Unsupported construction and private field access fail before backend emission.
6. Stage0 compilation, fixed point, complete CI and release smoke. New syntax remains in fixtures until a published stage0 supports it; no new extern-type standard resource is exposed to old stage0 as a shortcut.

The bounded helper type-query protocol, declaration syntax and compiler identity pipeline satisfy this sequence for module compilation. Standalone commands and source-only query/LSP analysis do not schedule Go metadata resolution. The complete plan requires the later B/C/D acceptance gates in addition to B0.
