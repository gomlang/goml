# syntax

A pure GoML lossless syntax tree library inspired by [Rowan](https://docs.rs/rowan/latest/rowan/). It provides immutable green nodes, contextual red views, checked parser construction, bounded interning, typed AST adapters and persistent tree edits. It has no native adapter or dependency outside the GoML standard library.

The module is `ecosystem::syntax`. The separate `consumer::syntax` module resolves version `0.1.0` and implements a nested configuration language with comments, whitespace, error recovery, typed values and trivia-preserving rewrites.

## Constructing a tree

Kinds are caller-defined `u32` values. The library assigns no meaning to a kind: whitespace, comments, malformed tokens and missing zero-width tokens can all be retained.

```goml
use ecosystem::syntax;

fn example() -> Result[syntax::SyntaxNode, syntax::Error] {
    let builder = syntax::GreenBuilder::new();
    builder.start_node(0)?;
    let expression = builder.checkpoint()?;
    builder.token(1, "name")?;
    builder.start_node_at(expression, 10)?;
    builder.token(2, " = ")?;
    builder.token(3, "\"value😀\"")?;
    let _ = builder.finish_node()?;
    let _ = builder.finish_node()?;
    Result::Ok(syntax::SyntaxNode::new_root(builder.finish()?))
}
```

`start_node_at` moves the current frame's children at and after a checkpoint into a new open node. This supports parsers which recognize a construct after consuming its left operand. A checkpoint belongs to one builder and one specific open frame; another builder, a closed/replaced frame, or a position past the current children produces `InvalidCheckpoint`. A checkpoint is a child-list position, not a stable edit anchor. Reusing a checkpoint after wrapping refers to that numeric position in the current list.

`element` inserts an existing immutable green node or token. `finish_node` closes exactly one frame and returns its immutable node; `finish` requires exactly one closed root node and seals the builder. Failed operations leave its valid construction state available for recovery. Cache hit/miss and eviction state can still change when token construction precedes a rejected insertion. The builder rejects orphan tokens, multiple roots, unfinished roots and all mutation after `finish`.

## Green values and ownership

`GreenToken::new(kind, text)` validates UTF-8 and copies the text into owned string storage. `GreenNode::new(kind, children)` copies the child slots into a `FrozenVec`. Descendants use private references that the library never mutates. Retaining a token does not retain an unrelated original input substring. `children().to_vec()` returns an independent outer container; modifying it cannot change the node.

Every node caches its logical byte length, element count, depth, maximum descendant fanout, child-end offsets and a structural fingerprint. All logical occurrences count, including repeated references to the same subtree. Cached bounds allow a parent to reject a too-large imported subtree without traversing it. Green `PartialEq`/`Eq` are structural; `Hash` uses the fingerprint and equality verifies actual contents on collisions. `same_allocation` explicitly tests sharing. Fingerprints are deterministic implementation details, not cryptographic hashes or a stable serialization format.

Green and red values are immutable snapshots. Plain assignment preserves a snapshot in constant time; independently constructed red roots have distinct occurrence identities, even when they share one green root. Green values and red views can be read and edited concurrently. There is no mutation or detach requirement before editing.

Builders and iterators have mutable cursors: copying them shares that state. They require caller synchronization if used concurrently. Constructing a new iterator creates a separate cursor. A checkpoint retains its builder state; a red child retains its ancestors and therefore its complete green root.

## API

| Area | Operations |
| --- | --- |
| Green construction | `GreenToken::new` / `with_limits`; `GreenNode::new` / `with_limits`; `GreenBuilder::new` / `with_options`, `start_node`, `checkpoint`, `start_node_at`, `token`, `element`, `finish_node`, `finish`, `open_depth` |
| Green inspection | `kind`, `text`, `len_bytes`, `element_count`, `depth`, `max_children`, `children`, `child`, `child_count`, `child_offset`, `fingerprint`, `same_allocation`, `structural_eq`, `validate_limits` as applicable |
| Red navigation | `SyntaxNode::new_root`, `green`, `kind`, `range`, `text`, `index`, `parent`, `root`, `same_tree`, `same_occurrence`, `child_at`, `first_child` / `last_child`, `first_child_or_token` / `last_child_or_token`, `next_sibling` / `prev_sibling` |
| Element/token navigation | `as_node`, `as_token`, `element`, `next_sibling_or_token`, `prev_sibling_or_token`, `first_token`, `last_token`, `next_token`, `prev_token`, `ancestors` |
| Iteration | `children`, `children_with_tokens`, `descendants`, `descendants_with_tokens`, `tokens`, `walk`; standard `Iterator` implementations work with `std::iter` |
| Text queries | Checked `TextRange::new`, `start`, `end`, `len`, containment and intersection; `token_at_offset`, `text_slice`, `covering_element` |
| Persistent editing | `GreenNode::splice_children`, red `replace_with`, `SyntaxElement::replace_in` |
| Typed AST | Consumer implementations of `AstNode::cast` / `syntax`; generic `ast_children[T: AstNode]` |
| Interning | `NodeCache::new`, `token`, `node`, `intern`, `statistics`, `clear` |

All green and red element types implement `ToString`, yielding their lossless text. Red `PartialEq`/`Eq` and `Hash` identify the same occurrence in the same red root; repeated shared or zero-width subtrees remain distinguishable by their ancestor/index path.

`children` excludes tokens. Both descendant iterators include the receiver; `walk` emits balanced `Enter`/`Leave` events for every node and token. Calling `skip_subtree` immediately after an `Enter` skips that element's descendants but still emits its `Leave`. Iterators are fused. `SyntaxNode::ancestors` includes the receiver; `SyntaxElement::ancestors` starts with its parent. Token navigation includes zero-width tokens and skips empty nodes. Iterating a subtree stays inside it; `next_token`/`prev_token` navigate the full containing root.

## Byte offsets and edits

Ranges are half-open UTF-8 byte ranges. EOF is a valid boundary. `TextRange` checks nonnegative ordered bounds without knowing a tree; text queries additionally check containment and UTF-8 character boundaries. Interior bytes of a scalar produce `Utf8Boundary`, never a substring panic.

`token_at_offset` returns `None`, `Single(token)` or `Between(left, right)`. At a shared token boundary it returns both nonempty neighbors. At the beginning/end of nonempty text it returns the one adjacent nonempty token. Empty nodes and zero-width tokens do not claim an offset, so an entirely empty tree returns `None`. Zero-width elements remain visible through structural traversal. Positions inside an ASCII token return `Single`; positions inside a multibyte scalar are errors.

`covering_element` returns the deepest first child that contains the entire range. Empty ranges deliberately prefer the leftmost containing child, including zero-width children. `text_slice` accepts absolute ranges relative to the containing root, including when called on a subtree.

`replace_with` rebuilds the ancestors and returns a new green **root**, sharing unaffected siblings. A node/token's typed method replaces its own category; the general `SyntaxElement` method can change a child from node to token or vice versa. A root replacement must be a node. The old view and its offsets remain valid for the old snapshot. Create a new red root to navigate the replacement. `replace_in` additionally requires the supplied root to be the exact red root that owns the target and reports `ForeignTree` otherwise. `splice_children` uses child-index half-open ranges and returns the changed node itself.

## Cache, limits and costs

`NodeCache::new(capacity, max_bytes, max_elements)` creates a shared, synchronized, FIFO interner. Capacity must be in `0..=4096`; the other budgets are nonnegative. Zero capacity disables retention. Statistics expose retained entries/bytes/elements and cumulative hits/misses/evictions; `clear` releases cached roots while preserving those counters and existing snapshots. Oversized entries are returned without retention. Logical bytes/elements of retained entries are summed, including shared descendants, so the budget conservatively overcounts sharing and also bounds empty syntax trees. These counters are not precise Go heap-byte measurements.

Cache lookup scans the bounded entries, compares fingerprints, then verifies structural equality. This favors a simple inspectable memory bound over a large unbounded hash-consing table. A channel gate serializes complete cache operations; callbacks are not executed under it. Cache aliases are safe for concurrent use. Construction may allocate a temporary candidate even on a cache hit.

Default limits are 64 MiB of logical UTF-8 text, 1,000,000 logical node/token occurrences, depth 4096 and 1,000,000 children in any node. `with_limits`/`with_options` and editing methods accept custom `Limits`. Depth includes both nodes and tokens: a node containing one token has depth two. Zero-byte and zero-child budgets are valid, but depth and element count must be positive. Aggregate counts use checked arithmetic, including enormous trees built through shared doubling. Builders also bound pending open construction, rather than waiting until the root is finalized. Limits apply to imported descendants as well as freshly constructed nodes.

For height `H`, maximum fanout `W`, logical elements `E` and text bytes `B`:

- Snapshots and cached summaries cost `O(1)`. Child lookup, child offsets, direct parent/sibling navigation and red view construction cost `O(1)`.
- Token-at-offset descends binary-searched child ends in `O(H log(W + 1))`; it does not scan or flatten the text.
- Creating a node costs `O(number of children)`, independent of descendant size. Replacing a leaf costs the sum of child counts on its ancestor path, plus new token text.
- Lossless text materialization, traversal and structural equality visit logical contents. Their iterative worklists avoid recursion; text/equality/edge-token searches can retain up to `O(E)` pending work in wide trees. Ordinary red traversal retains `O(H)` ancestor state. Green equality can skip shared subtrees.
- Text slicing scans the subtree's tokens and allocates only the selected string. Covering-element queries scan children at each selected depth. These APIs are not a rope-backed text index.
- Red equality and hashing use the occurrence path and cost `O(H)` in the worst case. Walking a subtree normally compares its boundary cheaply; repeated zero-width/shared shapes can increase those comparison costs.

Explicitly raising limits can describe enormous logical trees in little memory. Full text materialization or traversal still needs resources proportional to their logical contents. Allocation exhaustion and panics from application code are not converted into errors. There is no cancellation protocol for tree operations.

This package does not supply a grammar, lexer, incremental parser/reparse scheduler, mutation-in-place red nodes, persistent edit-position tracking, grapheme/UTF-16 line indexes, serialization format or compiler CST migration. The independent consumer shows how to layer parsing and domain diagnostics over the tree; `rope`, `logos` and `incremental` remain separate ecosystem libraries.

## Verification

Run from the repository root:

```sh
python3 ecosystem/verify.py syntax
```

The verifier checks formatting, 18 library black-box tests, three versioned consumer tests, fresh and cached consumer builds, a runnable example, the independent Python oracle and all 18 library tests under Go's race detector. No CI integration is required.

The oracle compares 320 tree models, 3,840 persistent edits and 240 configuration rewrites. It independently computes lossless text, full preorder paths/kinds/ranges, tree summaries, all selected byte offsets, Unicode slicing failures, changed trees and retained snapshots. Configuration cases exercise Unicode keys/strings, escapes, nested sections, comments, CRLF and rejected replacement literals. Unit tests add malformed builder states, foreign/stale checkpoints, cache budgets/eviction, duplicate shared zero-width occurrences, error recovery, typed AST trait bounds, 10,000-deep iterative operations, aggregate overflow without enormous allocations, and concurrent interning/read/edit workloads.
