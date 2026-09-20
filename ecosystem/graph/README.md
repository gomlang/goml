# ecosystem::graph

A generic mutable multigraph library implemented in GoML. The module supports
directed and undirected graphs with arbitrary node and edge values, parallel
edges, self-loops, stable handles, graph transformations and weighted algorithms.

```gom
use ecosystem::graph;

fn route_cost() -> Result[Option[i64], graph::Error] {
    let roads: graph::Graph[string, i64] = graph::Graph::new(false);
    let a = roads.add_node("a");
    let b = roads.add_node("b");
    let _ = roads.add_edge(a, b, 7)?;
    let paths = graph::dijkstra(roads, a, |distance| distance)?;
    Result::Ok(paths.distance_to(b))
}
```

## Graph operations

- `Graph[N, E]::new(directed)`, `is_directed`, `node_count`, `edge_count`.
- `add_node`, `add_edge`, `node`, `edge`, `set_node`, `set_edge`.
- `remove_node` also removes every incident edge; `remove_edge` and `clear`.
- `contains_node`, `contains_edge`, `nodes`, `edges`, `neighbors`,
  `predecessors`, `edges_between`.
- `map[A, B](node_map, edge_map)` and `copy` create an independent graph and
  return `Copy` with old-to-new node and edge handle maps.

`NodeId` and `EdgeId` are opaque, hashable, graph-specific handles. Their `index`
is useful for ordering and diagnostics, but is not their complete identity.
Passing another graph's handle returns `WrongGraph`; removed handles return
`InvalidNode` or `InvalidEdge`. Slots are never reused, including after `clear`.
Deletion releases labels but retains slot metadata. Copying into a fresh graph
compacts slots and changes every handle's identity.

Graphs have shared mutable storage: assigning or passing a graph preserves its
identity and mutations. `copy` duplicates topology and label values; labels that
contain references still share their referenced data. `nodes`, `edges` and
adjacency methods return independent outer vectors. Undirected adjacency emits
a self-loop once and keeps each parallel edge. Directed adjacency follows
insertion order; undirected adjacency lists outgoing edges before incoming ones.

## Algorithms

| API | Behavior |
| --- | --- |
| `bfs`, `dfs` | Iterative reachable-node traversal in deterministic order |
| `unweighted_path` | Minimum-edge path, including both endpoints; `None` when unreachable |
| `topological_sort` | Directed Kahn ordering; detects cycles including self-loops |
| `weak_components` | Components ignoring directed edge orientation |
| `strong_components` | Iterative Kosaraju SCCs; ordinary components for undirected graphs |
| `dijkstra` | Nonnegative `i64` weights, heap frontier, all reachable distances |
| `astar` | Targeted path, cached nonnegative heuristic, reopening of improved nodes |
| `bellman_ford` | Signed weights and reachable negative-cycle detection |
| `minimum_spanning_forest` | Undirected Kruskal forest; negative weights supported |
| `UnionFind` | Checked indices, path compression, union by size, incremental additions |

`Paths` exposes `start`, `distance_to`, `path_to`, `route_to`, and `reachable`.
`Route` contains node handles, exact edge handles and total cost, so parallel
edges are unambiguous. Unreachable targets return `None`; a source-to-source
route has one node, no edges and zero cost. Results describe the graph at search
time; later mutation can invalidate returned handles.

Weight callbacks are evaluated once per live edge. Dijkstra and A* reject
negative weights even in disconnected components. A* evaluates the heuristic
once per live node, rejects negative values and requires zero at the target.
The caller must supply an admissible heuristic for optimality; inconsistent
admissible heuristics are supported through reopening. Bellman–Ford ignores
unreachable negative cycles. In an undirected graph, a reachable negative edge
forms a negative closed walk and is reported as a negative cycle.

All cost accumulation is checked: intermediate overflow returns `Overflow`,
including an explored candidate that would not improve the final path. No
integer value is reserved as an infinity sentinel. `Forest` reports selected
edges, total cost and component count, including isolated live nodes; empty
graphs have zero components. Equal-weight forest edges are ordered by edge slot.

## Complexity and constraints

Let S be allocated node/edge slots (including deleted slots), V live nodes, and
E live edges. Enumeration scans the relevant slot array. Traversals and SCCs
take O(S + V + E); Dijkstra takes O(S + (V + E) log(V + E)) with a lazy heap;
Bellman–Ford takes O(S + VE); Kruskal takes O(S + E log E). A* can revisit nodes
with inconsistent heuristics. Adjacency lookup copies O(degree) values. Edge
deletion scans its endpoint adjacency vectors, and deleting a high-degree node
can be quadratic in its incident edge count.

Callbacks must not mutate the graph being processed. Shared storage is not
synchronized for concurrent access. Costs are `i64`; floating weights, graph
serialization, flow algorithms and dynamic shortest-path maintenance are outside
this API. Traversal and SCC implementations do not recurse on graph depth.

## Validation

The public API suite checks mutations, cross-graph identity, stale handles,
parallel edges, loops, generic maps, paths, overflow, negative cycles and forest
behavior. Independent checks include all 64 loop-free directed three-node
graphs, 80 deterministic weighted graphs compared against Floyd–Warshall from
every source, 32 complete four-node graphs compared against exhaustive spanning
tree subsets, A* reopening and a 3,000-node chain/cycle.

`../consumers/graph` uses normal versioned registry dependencies and downstream
label types. Run `python3 ecosystem/verify.py graph` from the repository root to
check both modules, build a consumer twice, verify cached artifact stability,
and run the executable.
