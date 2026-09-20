# Templates for GoML

A compiled template engine with expressions, lexical scopes, conditions,
filtered loops, includes, inheritance, HTML escaping and extensible functions
and filters. Parsing and evaluation are implemented in GoML; Unicode and math
operations use the standard library. Each render has its own scopes, operation
budget and output buffer. The template AST can be reused.

The syntax follows familiar [Jinja template conventions](https://jinja.palletsprojects.com/en/stable/templates/).
The supported dialect and its differences are specified below. Reference tests
cover the shared syntax; they do not establish full Jinja compatibility.

## Example and public API

```gom
use ecosystem::template;
use std::serde::{Serialize};

#[derive(Serialize)]
struct Page {
    title: string,
    names: Vec[string],
}

fn page() -> Result[string, template::Error] {
    let engine = template::Engine::new(template::Options::standard())?;
    engine.add("page", "<h1>{{title}}</h1><ul>{% for name in names %}<li>{{name}}</li>{% else %}<li>Empty</li>{% endfor %}</ul>")?;
    let context = template::from_serializable(Page {
        title: "People <2026>",
        names: Vec::from_array(["Alice", "中文 & GoML"]),
    })?;
    engine.render("page", context)
}
```

| API | Behavior |
| --- | --- |
| `Engine::new(Options)` | Validate limits and create an engine |
| `add(name, source)` | Compile and cache; replace an existing name only after successful compilation |
| `contains(name)`, `remove(name)` | Inspect/invalidate cached templates; remove reports whether one existed |
| `render(name, context)` | Load and render a named template |
| `render_string(source, context)` | Compile and render an unnamed source without adding it to the cache |
| `compile(name, source)`, `compile_with_limits(name, source, limits)` | Create a reusable `Template` with private AST |
| `render_template(template, context)` | Render an already compiled template |
| `Template::name()` | Return its diagnostic name |
| `register_function(name, Function)` | Register/replace `(Vec[Value]) -> Result[Value, string]` |
| `register_filter(name, Filter)` | Register/replace `(Value, Vec[Value]) -> Result[Value, string]` |
| `set_loader(Loader)` | Register `(string) -> Result[Option[string], string]` for cache misses |
| `escape_html(string)` | Escape an individual string under the standard output limit |

A loader returns `None` for a missing template and `Err` for an actual loading
failure. Successfully compiled loads are cached by name. Invalidate them with
`remove` or replace them with `add` after a source change. Loading is supplied by
the application; the engine does not implicitly read files or resolve paths.
Names passed to an application filesystem loader must be resolved according to
that application's intended template root.

Registrations can capture application state and override builtin names except
`super`. Unknown functions/filters and callback failures return errors at their
call site. Callback results are validated and copied before evaluation resumes.
Callbacks execute synchronously; the interpreter's operation budget does not
interrupt application callback code. The engine does not synchronize cache,
loader or registration mutation; callers coordinate those mutations.

Precompiled templates retain the parsing limits used to create them. Rendering
uses the engine's execution limits. Each render snapshots its context, including
nested lists and objects, so template assignments do not modify caller data.

## Context values

`Value` has `Undefined(string)`, `Null`, `Bool`, `Int(i64)`, `Float(f64)`,
`Text(string)`, `Safe(string)`, `List(Vec[Value])` and
`Object(Vec[(string, Value)])` variants. Objects preserve insertion order and
reject duplicate names during context validation. `get(name)` returns a value
or an `Undefined` path. `as_text()` accepts both ordinary and safe strings;
`as_int()` accepts integers.

`Value::from_json(std::json::Value)` converts JSON data. `from_serializable[T]`
uses the standard Serde-to-JSON value convention, then converts it. These
convenience conversions check finite numbers, signed 64-bit integer range,
depth 128, collection lengths 100,000 and total visited values 1,000,000.
Direct `Value` contexts use the engine's configured limits when rendered.

Output displays integers and finite floats as numbers, booleans as `true` or
`false`, null as an empty string, and lists/objects as compact JSON. Safe strings
contain explicitly trusted HTML. Collections are written through a bounded JSON
emitter instead of first creating one unbounded serialized string.

By default, using an undefined value in output, a condition, a loop or a
comparison fails. Missing attributes/indices can propagate through further
lookups so `missing.deep|default('fallback')` works. `is defined` and
`is undefined` inspect this state without raising. With
`strict_undefined: false`, undefined output is empty, its truth value is false,
and iteration is empty. Invalid operations on other value types still fail.

## Expressions

`{{ expression }}` prints a value. Names may contain Unicode letters and
numbers, starting with a letter or underscore. String literals accept either
quote, ordinary backslash escapes and JSON-style Unicode escapes. Decimal
integer/float literals, booleans, null/none, list/tuple literals and object
literals with evaluated string keys are supported.

Expressions include:

- Attribute/index access: `person.name`, `person['name']`, `items[-1]`, `pair.0`.
- Slicing: `items[start:stop:step]`, including omitted indices and negative steps.
  String indices/slices count Unicode scalar values.
- Arithmetic: `+`, `-`, `*`, `/`, `//`, `%`, `**`; unary `+` and `-`.
- Concatenation: `~` stringifies operands; `+` also joins two strings or two
  lists; strings and lists can be repeated with an integer multiplier.
- Comparisons: `==`, `!=`, `<`, `<=`, `>`, `>=`, `in`, `not in`, including
  short-circuit chains such as `0 < age < 120`.
- Logical `and`, `or`, `not`. `and`/`or` return the selected operand and skip
  unnecessary evaluation. Membership in an object checks its keys.
- Conditional expressions: `yes if condition else no`.
- Named calls and filter chains: `range(0, 10, 2)`, `names|sort|join(', ')`.
- Tests: `defined`, `undefined`, `none`/`null`, `number`, `integer`, `string`,
  `boolean`, `sequence`, `mapping`, `odd`, `even`; `is not` negates a test.

From low to high, precedence is conditional, `or`, `and`, comparisons, `~`,
addition/subtraction, multiplication/division/remainder, unary numeric operators,
power, then access/calls/filters. `not` includes comparisons in its operand.
Power associates to the right. Parenthesize a whole arithmetic expression
before filtering it: `((xs + ys) * 2)|join(',')`.

Integer arithmetic checks overflow. Integer floor division/remainder follow the
divisor's sign. Division produces a float; mixed numeric arithmetic uses f64 and
rejects nonfinite results. Float remainder is defined by `a - floor(a / b) * b`.
Integer/float comparisons inspect the floating representation so large integers
are not rounded to f64 merely to compare them. String ordering is case-sensitive
Unicode/UTF-8 lexical order. Structural list/object equality is bounded by the
operation/depth budget.

## Statements and scope

| Construct | Syntax and behavior |
| --- | --- |
| Comments | `{# ignored #}` |
| Raw text | `{% raw %}...{% endraw %}` |
| Assignment | `{% set name = expression %}` assigns in the current scope |
| Captured output | `{% set name %}...{% endset %}` captures rendered text; marked safe when autoescaping is active |
| Conditions | `{% if condition %}...{% elif other %}...{% else %}...{% endif %}` |
| Loops | `{% for item in values if condition %}...{% else %}...{% endfor %}`; filter and else are optional |
| Unpacking | `{% for key, value in data|items %}...{% endfor %}` |
| Local scope | `{% with a=expression, b=other %}...{% endwith %}`; all right sides see the outer scope |
| Escape scope | `{% autoescape false %}...{% endautoescape %}` |
| Include | `{% include name %}` or a list of candidate names; first existing template wins |
| Include context | `include name with expression`, optionally followed by `only`; standard `with context` / `without context` also work |
| Optional include | `include name ignore missing` suppresses missing-template errors, including an exhausted fallback list |
| Inheritance | `{% extends 'base' %}` and `{% block name %}...{% endblock name %}` |

Objects iterate over keys; strings iterate over scalars. Multiple loop variables
require each item to unpack to the corresponding length. Each accepted loop
iteration has a fresh scope. Loop and loop-else assignments stay inside that
scope. Conditions share their enclosing scope. Includes see the current context
unless `only`/`without context` is used, but their assignments remain local.

Filtered loops evaluate the filter before calculating metadata. `loop` provides
`index`, `index0`, `revindex`, `revindex0`, `length`, `first`, `last` and the outer
loop as `parent`. Empty or fully filtered collections run the optional else body.

An extending template must put `extends` before content and can then have only
blocks, assignments, captures and whitespace at the top level. This prevents
silently discarded child output. Parent names may be expressions. Multilevel
inheritance and nested blocks are supported; `super()` renders the next block
implementation. Block assignments are local. Duplicate blocks, mismatched end
names, invalid `super()` calls and include/inheritance cycles return errors.

Whitespace is otherwise preserved, including trailing newlines. `-` immediately
inside an opening or closing delimiter trims adjacent Unicode whitespace:
`{{- value -}}`, `{%- if ready -%}` and `{#- comment -#}`. Raw blocks honor trim
markers at their own boundaries. Delimiters inside quoted expressions and object
literals are recognized without prematurely closing the tag.

## Builtin functions and filters

Functions are `range(stop)`, `range(start, stop[, step])`, `len`, `str`, `bool`,
`int`, `float`, `list`, `min`, `max`, plus inheritance's `super()`.

| Filter family | Filters |
| --- | --- |
| Missing values | `default([fallback[, boolean]])`, alias `d` |
| HTML | `escape` / `e`, `safe`, `forceescape` |
| Text | `upper`, `lower`, `trim`, `capitalize`, `length` / `count`, `replace(old, new[, count])`, `urlencode` |
| Collections | `join([separator])`, `first`, `last`, `reverse`, `list`, `sort([reverse])`, `unique`, `batch(size[, fill])` |
| Mappings | `items`, `keys`, `values`, `map('attribute.path')`, `selectattr('attribute.path')`, `rejectattr('attribute.path')` |
| Selection | `select`, `reject`, using truth values |
| Numbers | `int`, `float`, `abs`, `round`, `sum`, `min`, `max` |
| JSON | `tojson`, escaping HTML-sensitive characters and returning safe JSON text |

Filters check arity and argument types. `default` substitutes only undefined
values unless its second argument is true. Numeric conversions are strict;
invalid text fails instead of silently becoming zero. `round` uses the standard
library's rounding to nearest, with halves away from zero. `sort` is stable and
case-sensitive; `unique` retains the first structurally equal value. Attribute
filters use positional path strings. Casing follows `std::unicode`, including
its Unicode version and simple-case behavior.

Autoescaping is enabled by default. Escaping covers `&`, `<`, `>`, `"` and `'`.
`Safe` values and the `safe` filter are explicit trust decisions. Escape avoids
re-escaping safe values; `forceescape` always escapes. Concatenation and joins
preserve safe fragments while escaping ordinary fragments when appropriate.
The `~` operator respects a disabled autoescape scope; safe-string `+` retains
safe-string composition rules. Case/trim filters preserve safe status; `replace`
returns ordinary text. `tojson` is useful for JSON text inside script content;
normal quoted HTML attributes still require the corresponding HTML escaping.

The dialect does not include macros/imports, arbitrary object method calls,
keyword-argument calls, recursive-loop syntax, loop-control extensions, custom
tag delimiters or Jinja's full filter/test catalog. Unsupported syntax produces
a diagnostic. Application functions and filters provide explicit extension
points. Boolean/null/collection display, power association, strict conversions,
attribute-filter arguments and some numeric/string details differ from Jinja.

## Diagnostics, limits and verification

`Error` exposes `kind`, `location`, `message` and runtime include `trace`.
Locations contain template name, byte offset and one-based line/Unicode-scalar
column. Callback failures retain the template call site. Included runtime
failures retain their original location and add caller locations.

`Options::standard()` enables strict undefined checking and autoescaping, with:

| Limit | Default |
| --- | ---: |
| Source bytes per template | 1 MiB |
| Syntax work, token/chunk bound (`max_nodes`) | 100,000 |
| Parsing/evaluation depth | 128 |
| Rendering operations, including context validation, comparisons and iterations | 1,000,000 |
| Output/capture/intermediate string bytes | 8 MiB |
| Elements/entries in an individual collection | 100,000 |
| Simultaneously active templates, including inheritance and the initial template | 32 |

Limits are configurable nonnegative integers. Nested context copies detect
cycles by reaching the depth bound. Intermediate repetition, ranges, JSON
emission and output accumulation check their respective limits. These are
interpreter resource limits, not a process-wide memory cap or a deadline for
registered application code.

From the repository root:

```sh
just ecosystem-test template
```

There are 11 external library tests and 3 independently resolved consumer tests.
They cover grammar, scopes, generic Serde context conversion, captured callbacks,
inheritance, loaders, cache replacement, safety propagation, malformed inputs,
numeric boundaries, cycles and limits. The consumer is built and run separately;
a cached rebuild must leave generated package artifacts unchanged.

Native consumer tests compare all 1,367 shared-syntax cases with retained expected
results independently generated by checksum-pinned
[Jinja 3.1.6](https://pypi.org/project/Jinja2/3.1.6/) and
[MarkupSafe 3.0.2](https://pypi.org/project/MarkupSafe/3.0.2/). This includes 720
Unicode slice cases, generated arithmetic, exact mixed-number comparisons,
escaping/captures, filtered loops, includes and multilevel inheritance. [Fixture provenance](../consumers/template/tests/data/README.md) records the exact
reference distributions, digests, seed, and origin. Running the tests requires
GoML, with no Python interpreter or reference packages. Logs are retained under
`ecosystem/_artifact/verification/template/`.
