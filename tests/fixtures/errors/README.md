# errors fixtures

Small, self-contained `.go` snippets that trigger (or deliberately avoid
triggering) the errors agents/scanners. Each fixture should be minimal —
one bug pattern per file — so it's obvious which agent it exercises.

<!-- TODO: add one fixture per agent in plugin/agents/errors/, plus at
     least one ACCEPTABLE (false-positive) counter-example per agent -->
