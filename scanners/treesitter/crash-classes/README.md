# Tree-sitter crash-class detectors

Tree-sitter based detectors that parse a real Go syntax tree to target
specific reachable crash classes: index-out-of-range on unchecked slice
access, nil-map-write, integer overflow on untrusted input, and
`sync.WaitGroup` misuse.

Validated against confirmed issues in real Go projects, same discipline
as the reference CPython toolkit's crash-class detectors.

<!-- TODO: one detector per crash class -->
