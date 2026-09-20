# ecosystem::cli

Command-line parsing with explicit command schemas, typed conversion, subcommands
and two third-party compile-time derives. Argument parsing is deterministic by
default; process arguments and environment lookup are explicit entry points.

```goml
use ecosystem::cli;
use cli::{Args, ArgValue};

#[derive(ArgValue)]
enum Mode { Fast, Safe }

#[derive(Args)]
#[command(name = "serve", about = "Serve files", version = "1.0")]
struct Options {
    #[arg(short = "p", default = "8080", env = "PORT")]
    port: u16,
    #[arg(default = "safe")]
    mode: Mode,
    #[arg(short = "v", count)]
    verbose: isize,
    #[arg(positional)]
    files: Vec[string],
}

fn options(arguments: Vec[string]) -> Result[Options, cli::Error] {
    cli::parse(arguments)
}
```

## Explicit schemas

Build a schema with `Command::new`, `about`, `version`, `arg`, `subcommand` and
`require_subcommand`. `Argument::new` consumes a value, `Argument::flag` records a
boolean switch, and `Argument::counter` counts occurrences. The free functions
`command`, `argument`, `flag` and `counter` provide equivalent constructors.

Argument builders include `short`, `help`, `required`, `multiple`, `positional`,
`default_value`, `env`, `possible_values`, `conflicts_with` and `requires`.
Positional indices are contiguous from zero; a repeated positional must be last.
Required positionals precede optional positionals. Commands and arguments are
public records; `validate` checks the schema, including duplicate names/short
options, invalid positions and unknown conflict/dependency targets.

`Command::parse` takes arguments without the executable name. Supported forms
include `--name value`, `--name=value`, `-p80`, `-p 80`, short-flag clusters such
as `-vvq`, and `--` to terminate option processing. Values beginning with `--`
must use the equals form. Use `--` for positional values beginning with `-`.
Scalar duplicates are rejected; repeated values preserve input order. Flags do
not accept explicit values. `--help`/`-h` and `--version`/`-V` are reserved.

Subcommands have their own schemas. Parent options precede the subcommand, and
`Matches::subcommand` returns the child's matches. No global-option inheritance
or fuzzy subcommand matching is applied. Schema nesting is capped at 64.

## Values and environment

`Matches` provides `raw`, `value[T]`, `optional[T]`, `values[T]`, `flag`, `count`,
`contains` and `was_provided`. Typed getters use `ArgValue`; standard
implementations cover strings, booleans, all integer widths and f32/f64.
Integer conversions enforce destination bounds. Applications may implement
`ArgValue` for custom types.

Precedence is explicit arguments, then environment, then defaults.
`parse_with_env` accepts an injected lookup function, making tests independent
of the host environment. `parse_env` reads process arguments (skipping argv[0])
and environment variables. Generic `cli::parse`, `cli::parse_with_env` and
`cli::parse_env` additionally decode an `Args` type.

Environment values count as provided for conflict/dependency checks; defaults do
not. A requirement may be satisfied by a default. Flags and counters do not take
environment values, defaults or choice lists. Repeated value arguments receive a
single value from an environment variable; automatic separator splitting is not
performed.

## Derives

`#[derive(Args)]` supports structs, including generic structs. Plain fields are
required unless they have a default, booleans become optional flags, `Option[T]`
is optional, and `Vec[T]` is repeated. A count field must be `isize`. Derived
bounds apply `ArgValue` to the actual scalar or container-element type.

`#[command(...)]` accepts string `name`, `about`, and `version`. `#[arg(...)]`
accepts switches `positional`, `required`, `optional`, `multiple`, and `count`,
and string values `long`, `short`, `help`, `default`, `env`, `choices`, `conflicts`
and `requires`. The last three use `|`-separated names/values. Empty string
defaults are preserved. Field underscores become hyphens in default option
names. Positional indices follow field order. Unknown attributes produce
compile-time diagnostics. Explicit command schemas handle subcommands.

`#[derive(ArgValue)]` supports nonempty enums with unit variants. Variant names
default to ASCII lowercase with underscores replaced by hyphens. A variant may
have `#[value(name = "plain-text", alias = "text")]`. Duplicate wire names and
payload variants are rejected. Both derives support imported aliases and use
definition-site helper identities, so caller names such as `Command` do not
capture generated references.

`command_for[T]()` returns `TypedCommand[T]`; use its public `command` field for
validation and help. Keeping `T` in this return type avoids the current compiler's
erased-generic-function issue documented in `../FINDINGS.md`.

## Errors and validation

`Error` distinguishes `Help`, `Version`, `Usage`, `InvalidValue` and
`InvalidSchema`, with a printable message. Help and version requests return
before required-value checks. The library never exits the process or prints on
the caller's behalf.

```sh
python3 ecosystem/verify.py cli
```

Tests exercise explicit schemas, clusters, terminators, duplicate/missing values,
environment precedence, conflicts, subcommands, generated generic decoding,
custom conversion, enum aliases, empty defaults, integer overflow and stable
help text. The independent consumer verifies registry-exported derives, import
aliases and generated-name hygiene.
