These native tests retain the generated and mutated inputs from the former
standard-library integration fixture. Files in `known_answers/` freeze Go 1.26.0
`net/url` and `net/netip` reference results for 42568 distinct argument sets.

Each row contains the operation and all arguments, followed by tab-separated
results. String and byte arguments/results use lowercase hexadecimal bytes;
booleans and integers use text. A missing answer fails the test. Tables load
independently when first needed. Running the tests requires no Go FFI or external
reference implementation.
