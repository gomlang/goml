<a id="modules-packages-and-visibility"></a>

# Modules, packages, and visibility

[English](packages.md) | [简体中文](packages.zh-CN.md)

[GoML language guide](../goml.md)

Source-level package identities, imports, and visibility. Module manifests, workspaces, dependency resolution, and build layout are covered in [project tools](../tooling.md).

On this page: [Modules, packages and imports](#modules-packages-and-imports).

<a id="modules-packages-and-imports"></a>

## Modules, packages and imports

<a id="directory-packages"></a>

### Directory packages

Each directory containing a `.goml` file is a package. All source files in the same directory must declare the same package name:

```goml
package utils;
```

The canonical identity of a package is the module path plus the relative directory path. The package declaration name determines the local name when imported without aliases, but not the global identity of the package. The canonical identity of the root package is the module path.

A typical project could be:

```text
goml.toml
main.goml
utils/utils.goml
```

`utils/utils.goml`:

```goml
package utils;

pub fn message() -> string {
    "hello"
}
```

`main.goml`:

```goml
package main;

use alice::myapp::utils;

fn main() -> () {
    println(utils::message())
}
```

<a id="use"></a>

### `use`

`use` is a file-level declaration and must be placed after `package` and before any top-level items. Imports are not shared between different files.

Imports are not transitive: package A imports package B, which does not allow files using A to automatically see B. Each file must directly import each package it references.

```goml
use alice::http::client;
use alice::http::client as http_client;
```

The contextual `module` path head denotes the canonical path declared by the current project's `[module].path`. It may be used from any package in that module, including modules whose canonical path contains multiple segments:

```goml
use module::http::client;
use module::rendering::api::{Canvas, Render};
pub use module::model::Request;
```

For a module declared as `alice::myapp`, these paths resolve to `alice::myapp::http::client`, `alice::myapp::rendering::api`, and `alice::myapp::model::Request`. The marker applies only at the start of a `use` path followed by `::`; dependency paths and aliases named `module` elsewhere retain their ordinary meaning. External dependencies continue to use canonical paths such as `alice::http::client`.

After importing the package, access its public items through local aliases:

```goml
let request = http_client::new_request();
```

Braced imports bring selected public top-level items directly into the current file. Functions, constants, structs, enums, type aliases, and traits can be imported, and each item may be renamed independently:

```goml
use alice::rendering::api::{Canvas, Color as Paint, Render, render_to_string};

fn describe(canvas: Canvas) -> string {
    render_to_string(canvas)
}
```

The path before the braces is always a package path. Nested braced imports, `self`, glob imports with `*`, enum-variant imports, and subpackage aggregation are not supported. Imported names remain file-local, and normal top-level visibility rules apply.

The following special form also adds a trait from an already loaded package to the method calling scope:

```goml
use alice::rendering::api;
use api::Render;
```

Importing `Render` in a braced list also adds it to the method scope. Afterwards, the specific value can be written as `value.render()`. Even if the trait is not added to the method scope, you can still write a qualified call to `api::Render::render(value)`.

`pub use` re-exports a named package or public item from the current package without changing its identity:

```goml
pub use alice::http::client;
pub use client::{Client, Request, Response};
pub use client::Request as HttpRequest;
```

Downstream code can use `facade::HttpRequest`, import it directly, or bring a re-exported trait into method scope. The compiler resolves every such path back to the original declaration, so constructors, trait implementations, methods, constants, and functions behave exactly as they do at the original path. Re-exported package namespaces also remain usable, for example `facade::client::Request`. Re-export names in one package must be unique. Glob re-exports remain unsupported.

The following path models are not supported:

```text
mod child;
crate::x
self::x
super::x
::rooted::x
```

<a id="visibility"></a>

### Visibility

Top-level items are only visible within the package by default. `pub(module)` permits access from other packages belonging to the same module; `pub` permits access from external modules too. Structure fields and inherent impl methods are private by default and need their own visibility annotation for cross-package access:

```goml
pub struct Point {
    pub x: i32,
    pub y: i32,
}

pub fn origin() -> Point {
    Point { x: 0, y: 0 }
}

impl Point {
    pub fn sum(self) -> i32 {
        self.x + self.y
    }
}
```

All files in the same package can use private top-level items. Trait impl methods inherit the trait contract and cannot declare either `pub` or `pub(module)`.

```goml
pub(module) struct DeviceState {
    pub(module) cycles: u64,
}

pub(module) fn restore_clock(state: DeviceState) -> u64 {
    state.cycles
}
```

Module membership follows the canonical module identity from `goml.toml`, not a shared workspace or a common path prefix. Workspace members, path dependencies, replacements, and registry dependencies do not gain access to another module's internal items. Fields, newtype fields, inherent methods, and top-level declarations use the same three visibility levels.

`pub(module) use` can expose an imported item to sibling packages without exposing it externally. A `pub use` cannot promote a module-internal declaration to public visibility. Public signatures must not expose module-internal types; use a public type or reduce the signature's visibility. Public documentation hides module-internal declarations; documentation requested with private items includes them.

Module-internal visibility is available to production code, not only tests. To verify the external public API, place consumer tests in a separate module.
