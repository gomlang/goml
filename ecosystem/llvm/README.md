# LLVM

`ecosystem::llvm` provides typed GoML bindings to LLVM 18 for constructing,
parsing, verifying, optimizing and compiling LLVM IR. Its public API uses
`Context`, `Module`, `Type`, `Value`, `Function`, `Block`, `Builder` and `TargetMachine` handles,
operation enums and recoverable `Result` errors. Ordinary Go FFI connects these
types to a cgo adapter calling the LLVM C API. No compiler hooks or runtime
externs are added.

## Installation and consumption

The current adapter targets Linux amd64 and the LLVM 18 development installation:

- Headers: `/usr/lib/llvm-18/include`.
- Shared library: `/usr/lib/llvm-18/lib/libLLVM-18.so`.
- Verification tools: `/usr/lib/llvm-18/bin`.
- A C compiler, Go 1.26 or newer and enabled cgo are required.

The default include and linker paths are declared in `adapter/adapter.go`.
`CGO_CFLAGS` and `CGO_LDFLAGS` can supply paths for another LLVM 18 installation;
its shared library must also be discoverable by the runtime loader. Other LLVM
major versions require an adapter change and fresh verification. `version()`
reports the linked LLVM version. The binding has no external Go module
dependencies.

A GoML consumer declares the normal versioned dependency:

```toml
[dependencies]
"ecosystem::llvm" = "0.1.0"
```

The library declares its Go adapter, required cgo, and LLVM major in `[native]`.
The consumer needs only a minimal [go.mod](../consumers/llvm/go.mod); the driver
generates requirements and replacements pointing at the selected registry source.
It verifies cgo and LLVM 18 before compilation. Registry publication includes the
adapter sources but does not install the LLVM shared library or set linker paths.

## Building a function

```gom
use ecosystem::llvm::{Context, BinaryOp, OptimizationLevel, LlvmError};

fn compile_add(path: string) -> Result[string, LlvmError] {
    let context = Context::new();
    defer {
        let _ = context.close();
    };
    let module = context.module("example")?;
    let integer = context.int_type(64)?;
    let signature = context.function_type(
        integer,
        Vec::from_array([integer, integer]),
        false,
    )?;
    let function = module.add_function("add_numbers", signature)?;
    let entry = function.append_block("entry")?;
    let builder = context.builder()?;
    builder.position_at_end(entry)?;
    let sum = builder.binary(
        BinaryOp::Add,
        function.parameter(0)?,
        function.parameter(1)?,
        "sum",
    )?;
    builder.ret(sum)?;
    module.verify()?;
    module.optimize(OptimizationLevel::Default)?;
    module.emit_object(path, OptimizationLevel::Default)?;
    module.ir()
}
```

Link the resulting object with a C declaration such as
`int64_t add_numbers(int64_t, int64_t)`. The caller is responsible for choosing
types and signatures compatible with its ABI.

## Public API

| Area | Operations |
| --- | --- |
| Context | `new`, `module`, `builder`, `target_machine`, `parse_ir`, `parse_bitcode`, `close`, `is_closed` |
| Types | `void_type`, `int_type`, `float_type`, `double_type`, opaque `pointer_type`, `array_type`, literal `struct_type`, `function_type` |
| Type values | `same_as`, `ir`, `const_int`, `const_signed`, `const_float`, `zero`, `undef`, `const_aggregate` |
| Module | `add_function`, `function`, `ir`, `verify`, `clone`, `bitcode`, `target_triple`, `data_layout`, `optimize`, `run_passes`, `emit_object`, `emit_assembly`, `close` |
| Target discovery | `target_names`, `normalize_triple`, `TargetOptions::new`, `TargetOptions::native` |
| Target machine | `triple`, `cpu`, `features`, `name`, `data_layout`, `configure`, `optimize`, `run_passes`, file/memory emission, `close`, `is_closed` |
| ABI layout | `byte_order`, `pointer_bytes`, `pointer_int_type`, `type_layout`, `field_offset` |
| Function | `signature`, `parameter`, `append_block`, `blocks`, `name`, `same_as`, `entry_alloca`, `local`, `ir` |
| Block | `instructions`, `terminator`, `predecessors`, `successors`, `name`, `same_as` |
| Value | `value_type`, `name`, `set_name`, `same_as`, `operands`, `users`, `set_operand`, `replace_all_uses_with`, `erase`, `ir` |
| PHI | `is_phi`, `incoming`, `add_incoming`, `replace_incoming`, `set_incoming`, `remove_incoming` |
| Insertion points | `position_at_end`, `position_at_start`, `position_before`, `save_position`, `restore_position`, `clear_position` |
| Local variables | `Local::read`, `Local::write`, `Module::promote_locals` |
| Builder arithmetic | `binary`, `int_compare`, `float_compare`, `cast`, `select` |
| Builder control flow | `position_at_end`, `call`, `ret`, `ret_void`, `branch`, `conditional_branch`, `phi`, `unreachable` |
| Builder memory | `alloca`, `load`, `store`, `gep` |

All handle types expose `is_closed`. Operations that can fail return
`Result[..., LlvmError]`; errors distinguish `Argument`, `Closed`, `Stale`,
`Verification` and `Llvm`, with an accompanying diagnostic message.

Integer widths range from 1 to 65,536 bits. Array lengths are capped at
2,147,483,647 elements; pointer address spaces fit 24 bits. Floating types are
32-bit float and 64-bit double. Structs can be packed; function signatures can
be variadic. Arrays and structs accept checked aggregate constants.
`const_int(bits, sign_extend)` accepts a 64-bit pattern, truncating it for narrower
types and extending it for wider types. `const_signed(i64)` requests signed
extension. Arbitrary-width integer parsing is not exposed.

`BinaryOp` covers integer arithmetic, signed/unsigned division and remainder,
bitwise operations and shifts, plus floating arithmetic. `IntPredicate` includes
signed/unsigned comparisons; `FloatPredicate` includes ordered/unordered
comparisons. `CastOp` covers integer extension/truncation, numeric floating
conversions, pointer/integer conversion and checked equal-width bitcasts.

Each block accepts one terminator, and phi instructions must precede its non-phi
instructions. Builders support block-end insertion and explicit insertion points.
The module verifier checks the finished control-flow graph, SSA dominance and
remaining IR invariants.

LLVM 18 uses opaque pointers. `load` and `gep` therefore take explicit source
types. GEP checks aggregate traversal and constant i32 struct indices. The
caller remains responsible for actual pointer validity, memory layout, bounds,
alignment, ABI compatibility and LLVM poison/undefined-behavior rules.

## Constructing and editing SSA

`position_at_start(block)` inserts before the first non-PHI instruction, or at
the end of a block containing only PHIs. It permits adding PHIs after the rest
of the block has already been constructed. `position_before(instruction)`
selects an exact anchor. Ordinary instructions cannot be inserted before a PHI;
terminators can only be inserted at the end of an unterminated block.
`save_position()` returns an `InsertionPoint`, including an unset position;
`restore_position(point)` restores it in any builder of the same context.

`incoming()` returns ordered `(Value, Block)` pairs. `add_incoming` appends,
`replace_incoming` replaces the complete list, `set_incoming(index, value, block)`
changes one entry, and `remove_incoming(index)` removes one entry. Edits validate
types and function ownership before changing IR. Parallel edges from the same
predecessor are supported when they carry the same value. The verifier checks
edge multiplicity against the completed CFG; edits may temporarily leave an
incomplete PHI, including an empty list. Replacing or removing entries preserves
PHI handle identity, users, names, metadata and fast-math flags.

`module.functions()`, `function.blocks()` and `block.instructions()` return
snapshots in IR order. Predecessors and successors include repeated entries for
parallel edges. `block.terminator()` returns an optional instruction.
`value.operands()` distinguishes `Operand::Value`, `Operand::Block` and
`Operand::Function`; metadata and other special operands return an error.
`value.users()` accepts local instructions and parameters and returns distinct
instruction users in LLVM use-list order. `same_as` compares native identity.

`set_operand(index, value)` supports arithmetic, shifts, casts, comparisons,
alloca/load/store, return, select, freeze, conditional-branch conditions and
ordinary call arguments. It checks type and ownership. PHIs use their dedicated
editing methods; GEP indices, branch destinations, callees and other specialized
operands are not editable through this method. To change a terminator, erase it,
position the builder at block end and create its replacement, then repair PHIs.

`replace_all_uses_with(value)` rewrites uses of a local instruction or parameter
with a same-typed value from the same function or an eligible constant.
`erase()` only removes instructions without users. Erasure is explicit and may
remove side effects; it is not a dead-code analysis. Every alias of the erased
instruction becomes closed, builders anchored before it become unpositioned,
and saved positions anchored before it become closed. Other handles stay live.
These edits do not automatically repair dominance; call `module.verify()`
after completing the transformation.

For mutable frontend variables, `function.local(type, name)` allocates storage
in the entry block without moving existing builders. `local.write(builder, value)`
checks the variable type and emits a store; `local.read(builder, name)` emits a
load. Initialize every variable on all paths before reading it. The lower-level
`entry_alloca(type, name)` returns the storage pointer when address access is
needed. Entry allocation also works after the entry terminator exists.

`module.promote_locals()` runs `function(sroa,mem2reg)`, allowing LLVM to build
PHIs for eligible variables across branches and loops. It verifies before and
after transformation and invalidates module handles just like `run_passes`.
Reacquire functions, blocks and instructions through the traversal APIs afterward.
Address escape and unsupported memory uses can prevent promotion. See the
[LLVM mutable-variable tutorial](https://releases.llvm.org/18.1.8/docs/tutorial/MyFirstLanguageFrontend/LangImpl07.html)
and the executable [nested-loop example](../consumers/llvm/tests/ssa_execution_test.gom).

## Resource lifecycle and concurrency

Handle copies share native state. Native resources require explicit closure;
GoML garbage collection does not dispose LLVM contexts. Close the context in a
`defer` block for a scope containing its modules and builders.

- Closing a context disposes its builders, modules and target machines, invalidating every handle
  associated with that context.
- Closing a module invalidates its functions, blocks and instruction values.
  Builders positioned in it lose their insertion position but remain reusable.
- Closing a builder releases the builder alone. Context, module and builder
  closure is idempotent.
- Closing a target machine releases its machine and target-data layout. Closure
  is idempotent; modules and types created through that target remain valid.
- Types and context-owned constants remain valid until their context closes.
  Pure constant-folded builder results are context-owned. Inspected or folded
  constants referencing globals or block addresses retain module ownership.
  Constant graphs exceeding the inspection budget conservatively retain module ownership.

Calls on one context are serialized by a native mutex, including close. Separate
contexts can run independently. Cross-context operands, cross-module function
references and cross-function local operands are rejected. Serialization makes
shared handles safe to call concurrently; users still need to coordinate the
order of edits to a shared IR graph or builder insertion point.

Optimization invalidates all existing function, block and nonconstant value
handles belonging to the module, and clears its builders' insertion positions.
This occurs once a valid pass request reaches LLVM, including a request whose
pipeline later fails. Old handles report `Stale`; reacquire functions with
`module.function(name)` or construct fresh IR handles. Types and context-owned
constants remain usable. A closed or stale handle reports `is_closed() == true`.

## IR, passes and native output

`parse_ir` and `parse_bitcode` create modules in the receiving context. Parsing
does not substitute for `verify()`. `ir()` returns text; `bitcode()` verifies the
module and returns copied bytes. `clone()` creates an independently disposable
module in the same context.

`OptimizationLevel::{None, Less, Default, Aggressive}` maps to LLVM levels
0, 1, 2 and 3. `optimize` uses the corresponding default pass pipeline;
`run_passes` accepts LLVM's textual pass-pipeline syntax. Modules are verified
before and after optimization. Invalid arguments and reported LLVM errors are
returned to GoML, but arbitrary LLVM pipelines and native code generation are
not an isolation boundary for hostile input or internal LLVM failures.

Object and assembly output verify the module, create a native target machine
with the host CPU/features and PIC relocation, then compile a module clone.
Emission preserves the original IR and its live handles. Output paths are
written directly and may overwrite existing files; parent directories must
exist. These original `Module` convenience methods use the host's target and
features. They now reject an existing incompatible module triple or data layout
instead of silently replacing it. Explicit target selection uses `TargetMachine`.

## Explicit targets and ABI layout

```goml
use ecosystem::llvm::{Context, Module, TargetOptions, OptimizationLevel, LlvmError};

fn compile_arm(context: Context, module: Module, path: string) -> Result[(), LlvmError] {
    let target = context.target_machine(TargetOptions::new("aarch64-unknown-linux-gnu"))?;
    defer { let _ = target.close(); };
    target.optimize(module, OptimizationLevel::Default)?;
    target.emit_object(module, path)
}
```

`TargetOptions::new(triple)` selects the backend's baseline CPU, no additional
features, default optimization, PIC relocation and the backend's default code
model. `TargetOptions::native()` explicitly uses the host triple, CPU and feature
list. Public fields configure `triple`, `cpu`, `features`, `optimization`,
`relocation` and `code_model`. Code-generation optimization and the IR pass level
are separate settings. Features use comma-separated `+feature,-feature` syntax.

`target_names()` lists available machine-code backends in sorted order. All
backends linked into LLVM are initialized once. `normalize_triple` uses LLVM's
normalization, including missing triple components; it does not check backend
availability or promise to canonicalize architecture aliases. Machine creation
checks availability and syntax. LLVM determines the meaning and availability of
individual CPU/features names, and may report unknown names to stderr.

`Relocation` supports `Default`, `Static` and `Pic`. `CodeModel` supports `Default`,
`Small` and `Large`; explicit Small/Large are accepted on x86-64 and AArch64,
and Small on i386. Other backends require Default in this binding. Unsupported
model/backend combinations return `Argument` before entering LLVM. A valid
configuration still requires IR, CPU features and operating-system ABI choices
compatible with that backend; native LLVM is not a process-isolation boundary.

`configure(module)` checks context ownership and existing target metadata,
then assigns the machine's normalized triple and data layout. Existing nonempty
triples must normalize to the same string, and existing nonempty layouts must
match the machine's layout string exactly. A mismatch returns `Argument` without
changing the module. Configuration alone preserves builders and value handles.
`TargetMachine::optimize` and `run_passes` verify and configure the module, pass
the target machine to LLVM's pass manager, and apply the same handle invalidation
rules as module optimization. Invalid pipeline syntax reported by LLVM can leave
the module configured and its previous handles stale.

`emit_object(module, path)` and `emit_assembly(module, path)` compile a clone and
leave source metadata and handles unchanged. One unconfigured module can thus
be emitted for multiple targets. Explicitly configured modules require a matching
machine. `object_bytes(module, max_bytes)` returns independent `Bytes`, while
`assembly_text(module, max_bytes)` validates UTF-8. The maximum copy size must be
between 1 byte and 1 GiB. It bounds the returned copy; LLVM builds its native
output buffer before that check, so it does not bound LLVM's working memory.

`type_layout(type)` returns `TypeLayout` with bit size, byte storage size,
ABI allocation size, ABI alignment and preferred alignment. `field_offset`
reports a checked struct member offset. `pointer_bytes(address_space)` and
`pointer_int_type(address_space)` use the selected target, including address
spaces whose pointer width differs from the default. `byte_order()` returns
LittleEndian or BigEndian. The returned pointer integer type belongs to the
context and survives target closure.

Layouts reject void/function/opaque/recursive unsized types, scalable vectors
including aggregate members, more than 64 nesting levels or 65,536 traversal
visits, and ABI sizes exceeding 1 TiB. Shared subtypes cannot bypass the nesting
limit. Layout queries use target ABI rules, so a struct's offsets and alignment
can differ between 32-bit and 64-bit targets. Cross-compilation emits objects;
linking or executing them still requires the destination platform's toolchain,
libraries and runtime.

## Scope and verification

This binding covers common IR construction and native ahead-of-time compilation.
It does not expose the full LLVM API: JIT/ORC execution, explicit ABI-name selection,
debug metadata, globals, named recursive structs, vectors, atomics, exception
handling, linkage/attribute configuration, arbitrary instruction mutation,
dominance/loop analysis, MemorySSA and a sealed-block SSA builder are not
implemented. Calling conventions use LLVM's default C convention.

From the repository root:

```sh
just ecosystem-test llvm
```

Nineteen GoML library tests, five consumer tests and twelve native adapter tests
cover errors, concurrency and resource lifetimes. A GoML consumer invokes installed LLVM 18 tools and
cc through `std::process`, verifies emitted IR/bitcode, assembles/disassembles it,
links unoptimized/optimized objects, assembly and an explicit portable x86-64
target, and compares 12,030 native
function results with independently computed arithmetic expectations. The shared
verifier handles fresh/cached builds and race-detector execution. Independent
`llvm-readobj`, `llvm-objdump` and `llvm-mc` checks identify, disassemble and
reassemble x86-64, i386, AArch64 and big-endian AArch64 output. Cross-target
objects are inspected, not executed. Tests also cover target mismatch rejection,
packed/unpacked ABI layouts, shared target emission, concurrent close, output
copy limits and the supported code-model/relocation combinations.

SSA regressions cover parallel edges, late PHIs, incoming-list edits, self edges,
metadata and alias preservation, invalid dominance, rejected edits, insertion
points, erasure and concurrent close. An additional native workload promotes
nested loops with multiple backedges, inspects the generated PHIs, and checks
390 results across original, promoted and optimized objects. It also exercises
late PHI creation, parallel edges, operand edits, use replacement and erasure.

The APIs follow LLVM 18's [target-machine C interface](https://github.com/llvm/llvm-project/blob/llvmorg-18.1.8/llvm/include/llvm-c/TargetMachine.h),
[target-data interface](https://github.com/llvm/llvm-project/blob/llvmorg-18.1.8/llvm/include/llvm-c/Target.h)
and [object emission tutorial](https://releases.llvm.org/18.1.8/docs/tutorial/MyFirstLanguageFrontend/LangImpl08.html).
