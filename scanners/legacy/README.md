# Legacy scanners

stdlib-only (`go/ast`, `go/parser`, `go/types`) scanners for the legacy
analysis dimensions: unchecked errors, lost error context, context misuse,
defer traps, interface nil semantics, and complexity.

These scanners require no external dependencies and run anywhere Go
itself runs. They report **candidates** — the agents in `plugin/agents/`
read the actual code to confirm or dismiss each one.

<!-- TODO: one .go file per dimension, e.g. unchecked_errors.go, defer_traps.go -->
