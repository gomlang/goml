# Import a standard-library allowlist

From this directory, with the repository toolchain built:

```sh
../../stage2/bin/goml bind-go bindings.json --dry-run
../../stage2/bin/goml bind-go bindings.json
../../stage2/bin/goml check
../../stage2/bin/goml run
```

The output is:

```text
3
42
true
```

The JSON selects only `strings.Count` and `strconv.Atoi`. Generation typechecks Go metadata without executing initializers or changing dependency manifests. Text conversion and error handling remain explicit in `main.gom`.

Review `bindings/generated.gom`, `native/generated.go` and `bindings.json.goml-bind.json`. Repeating generation preserves identical files. Editing either generated source or its ownership manifest causes regeneration to fail; put handwritten wrappers in a separate file.
