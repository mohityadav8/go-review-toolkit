# Concurrent slice append

**Category:** races
**Agent ID:** `slice-race`

## What this agent looks for

Flags `append` on a shared slice from multiple goroutines, which can silently drop writes.

## Detection strategy

<!-- TODO: describe the AST/tree-sitter pattern or heuristic this agent dispatches on -->

## Classification guidance

When this agent's scanner produces a candidate, classify it as one of:

- **FIX** - confirmed bug, should be fixed before merge
- **CONSIDER** - plausible risk, worth a second opinion but not blocking
- **POLICY** - intentional pattern that violates a general rule but is acceptable here by project convention
- **ACCEPTABLE** - false positive or a pattern that is safe in this specific context

<!-- TODO: add 2-3 worked examples showing FIX vs ACCEPTABLE for this specific check -->

## Known false positives

<!-- TODO: list patterns that commonly trigger this detector but are not actually bugs -->

## Related agents

<!-- TODO: cross-link agents in the same category or adjacent categories -->
