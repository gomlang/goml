# Go interoperability

GoML calls Go through explicit foreign declarations. The compiler checks the declared GoML types and trusts the binding author's signature. The final Go build checks native symbol existence, assignability, method sets and generic constraints. `goml check` and editor analysis do not load Go package metadata.

```goml
use std::ffi;

#[go_type("os", "File")]
extern type File;

#[go_ffi("os", "Open")]
extern fn open_file(name: ffi::String) -> (ffi::Ptr[File], ffi::Error);

#[go_method("Close")]
extern fn close_file(file: ffi::Ptr[File]) -> ffi::Error;
```

An external type is opaque. Its import path, Go name and declared type arguments identify it; GoML does not inspect its fields, normalize Go aliases, discover methods or derive interface structure. Bind functions and methods explicitly, or provide a small Go shim when a native API needs conversion. `#[go_interface]` remains available with an explicit raw type, trait signatures and method mappings; final Go compilation checks the generated adapter against the real interface.

Use a module-root `go.mod` for local Go shims and Go dependencies. Ordinary project builds disable Go workspace mode, select module mode from the presence of `go.mod`, and otherwise use the user's Go environment for dependency and toolchain behavior. Go library export and native SIMD builds retain readonly, offline module handling and require dependencies to be prepared beforehand. Unicode validity, ownership, mutability and concurrency guarantees remain explicit adapter contracts; successful Go compilation cannot prove them.

`--ffi-check required` and `--ffi-check off` remain accepted for command compatibility. Both use trusted declarations and neither invokes metadata inspection or emits an unverified-binding warning. Unused foreign declarations may not appear in generated Go and therefore need not be checked by the final build.

See the [language guide](../goml.md#go-ffi) for the complete boundary rules, the [file reading example](../../goml/testdata/ffi/file-read/README.md) for explicit pointer and method bindings, and [Go library export](../goml.md#exporting-a-go-library) for the reverse direction.
