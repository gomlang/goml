# B2 raw Go errors and explicit outcomes

`std::ffi::Error` is the Go universe `error` interface, represented by its named identity with an empty import path. It retains both the dynamic type and value. `nil_error()` constructs an interface nil; `error_is_nil(value)` compares the interface itself with nil. A typed nil pointer inside an error is non-nil and is preserved unchanged.

`NonNilError::from_error` returns `Option[NonNilError]`, rejecting a true nil interface. The wrapper's storage is private and `as_error()` returns the original interface. Standard package consumers now receive filtered public interfaces during checking, while code generation retains full standard library layouts. This prevents direct construction through a private standard field. `Result[T, NonNilError]` rejects `Err(nil_error())` by type checking.

`Outcome[T]` has public value and error fields and an explicit `new(value, error)` constructor. It preserves partial values alongside failures. `into_result_discarding_value_on_error()` explicitly opts into dropping the value on failure and returns `Result[T, NonNilError]`. No tuple, nil, error or Result conversion is implicit.

The driver regression verifies Go error aliases, cross-package/generic/closure transport, typed nil versus interface nil, explicit Error method calls, errors.Is across a wrapped chain, simultaneous partial value/error preservation, optional nil payloads, checked wrappers and private construction rejection. Public error_matches and NonNilError::matches delegate to Go errors.Is. NonNilError::message_bytes copies the original Error() bytes, and message validates UTF-8 before returning a GoML string.

B2 now has public message and matching adapters. The full acceptance audit and latest CI evidence are recorded in acceptance-b2.md. Raw text handling must follow B3, and the remaining B1 os.File fixtures must use these representations without suppressing errors or discarding partial reads.

`Outcome::from_raw_result` accepts Result[T, Error] and a caller-supplied failure value. It returns Result[Outcome[T], BoundaryError], using the success value with nil on Ok, or the supplied failure value with the original non-nil error on Err. Err(nil) returns BoundaryError::NilError with the message "Go FFI boundary rejected Err(nil)". Typed-nil interfaces remain non-nil errors. `from_result` applies the same check to Result[T, NonNilError]. `into_tuple` exposes both fields unchanged. Failure values follow ordinary eager argument evaluation.

The reverse-conversion regression passes the resulting pair to actual Go functions, checking the selected value, sentinel identity, success nilness and typed-nil payload. It also verifies the distinct contract-error branch for Err(nil), and runs again through cached artifacts.

A real method-binding fixture now checks os.Open, File.Read and explicit File.Close using an explicit MutSlice view. Closing an alias invalidates reads through the original pointer. Nil File.Close retains its Go error behavior. A Reader returning bytes with n > 0 and io.EOF verifies that both the mutated buffer and Outcome value/error survive, and that explicit Result conversion retains the same EOF error.

Message conversion does not catch panics or change typed-nil receiver behavior. Invalid UTF-8 returns Utf8Error without replacement, while message_bytes preserves all bytes in separate mutable storage. The matching implementation retains errors.Is behavior, including custom Is methods, chain traversal and nil matching; same message text does not imply identity.
