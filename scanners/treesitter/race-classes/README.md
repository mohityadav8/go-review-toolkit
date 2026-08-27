# Tree-sitter race-class detectors

Tree-sitter based detectors targeting data-race classes beyond what
`go test -race` catches statically: goroutine-leak detectors (fire-and-forget
goroutines with no cancellation path), shared mutable state without a lock,
and concurrent map/slice access.

<!-- TODO: one detector per race class -->
