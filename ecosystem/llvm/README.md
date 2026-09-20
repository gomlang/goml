# LLVM

`ecosystem::llvm` provides typed GoML bindings to LLVM 18 for constructing,
parsing, verifying, optimizing and compiling LLVM IR. Its public API uses
`Context`, `Module`, `Type`, `Value`, `Function`, `Block` and `Builder` handles,
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
| Context | `new`, `module`, `builder`, `parse_ir`, `parse_bitcode`, `close`, `is_closed` |
| Types | `void_type`, `int_type`, `float_type`, `double_type`, opaque `pointer_type`, `array_type`, literal `struct_type`, `function_type` |
| Type values | `same_as`, `ir`, `const_int`, `const_signed`, `const_float`, `zero`, `undef`, `const_aggregate` |
| Module | `add_function`, `function`, `ir`, `verify`, `clone`, `bitcode`, `optimize`, `run_passes`, `emit_object`, `emit_assembly`, `close` |
| Function | `signature`, `parameter`, `append_block`, `ir` |
| Value | `value_type`, `set_name`, `add_incoming`, `ir` |
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

Builders insert at the end of a block. Each block accepts one terminator, and
phi instructions must precede its non-phi instructions. Add phi predecessors
through `value.add_incoming(Vec[(Value, Block)])`; this checks matching types,
function ownership and duplicate predecessors. The module verifier checks the
finished control-flow graph, SSA dominance and remaining IR invariants.

LLVM 18 uses opaque pointers. `load` and `gep` therefore take explicit source
types. GEP checks aggregate traversal and constant i32 struct indices. The
caller remains responsible for actual pointer validity, memory layout, bounds,
alignment, ABI compatibility and LLVM poison/undefined-behavior rules.

## Resource lifecycle and concurrency

Handle copies share native state. Native resources require explicit closure;
GoML garbage collection does not dispose LLVM contexts. Close the context in a
`defer` block for a scope containing its modules and builders.

- Closing a context disposes its builders and modules, invalidating every handle
  associated with that context.
- Closing a module invalidates its functions, blocks and instruction values.
  Builders positioned in it lose their insertion position but remain reusable.
- Closing a builder releases the builder alone. Context, module and builder
  closure is idempotent.
- Types and context-owned constants remain valid until their context closes.
  Constant-folded builder results may also be context-owned.

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
exist. Output uses the build host's target and features, so it is not a
cross-compilation or portable CPU-baseline API.

## Scope and verification

This binding covers common IR construction and native ahead-of-time compilation.
It does not expose the full LLVM API: JIT/ORC execution, target selection,
debug metadata, globals, named recursive structs, vectors, atomics, exception
handling, linkage/attribute configuration and arbitrary instruction inspection
are not implemented. Calling conventions use LLVM's default C convention.

From the repository root:

```sh
python3 ecosystem/verify.py llvm
```

The verification entry point formats/checks the library and independent registry
consumer, runs their tests, checks fresh/cached builds, then runs `interop.py`
and `race.py`. Native tests and GoML tests exercise resource errors and shared
lifetimes. Interoperability checks use installed LLVM tools and a C compiler to
verify emitted IR/bitcode and execute linked object code.
