# Go Code Review Toolkit

A Claude Code plugin — and standalone CLI — for exploring, analyzing, and reviewing Go source code. It answers the question: **where are the bugs, race conditions, and maintenance risks in this Go codebase?**

Built for Go's specific concerns — goroutine lifecycle, channel discipline, error handling, context propagation, interface nil semantics — not generic linting. This is not `go vet` or `golangci-lint` with extra rules; it's a review campaign that reads code the way a senior Go reviewer would, then classifies every finding as **FIX / CONSIDER / POLICY / ACCEPTABLE**.

Works on any Go repository: standard library, Kubernetes-ecosystem projects, or your own codebase.

## Installation

### From the marketplace (recommended)

```bash
# Add the marketplace (one-time setup)
claude plugin marketplace add <org>/go-review-toolkit

# Install the plugin
claude plugin install go-review-toolkit@go-review-toolkit
```

### Direct install from GitHub

```bash
claude plugin install go-review-toolkit --source github:<org>/go-review-toolkit --path plugins/go-review-toolkit
```

### Without installing (try it first)

```bash
git clone https://github.com/<org>/go-review-toolkit.git
claude --plugin-dir go-review-toolkit/plugins/go-review-toolkit
```

### After upgrading the plugin

Run `/reload-plugins` before `explore` or `informed-explore`. A Claude Code session registers dispatchable agents when the plugin is loaded; agents added by a version installed mid-session stay invisible until you reload, and those commands then run silently with the older, smaller agent set.

## Quick Start

Navigate to a Go source checkout, then:

```bash
/go-review-toolkit:map              # Understand package/module structure
/go-review-toolkit:health           # Quick health dashboard
/go-review-toolkit:hotspots         # Crash-class + race + complexity detectors
/go-review-toolkit:explore          # Full exploration (all agents)
/go-review-toolkit:known-issues     # Regression check vs the known-bug catalog
/go-review-toolkit:reproduce        # Turn a static candidate into a reproduced failure
```

Start with `map` to understand the module/package graph, then `hotspots` to find the highest-impact bugs.

## What's Included

**26 analysis agents** covering goroutine lifecycle and leaks, channel discipline (nil-channel blocks, unbuffered send-without-receiver, close-of-closed-channel), data races beyond what the race detector catches statically, deadlock patterns (lock-ordering, self-deadlock via re-entrant `Lock()`), error handling (unchecked errors, lost error context, `errors.Is`/`As` misuse, sentinel-error anti-patterns), context propagation (context leaks, missing cancellation, `context.Value` misuse, ignored `ctx.Done()`), defer traps (defer-in-loop, defer with unevaluated arguments, panic-in-defer swallowing the original panic), interface nil semantics (typed-nil-in-interface, nil-pointer-satisfies-interface footguns), slice/map aliasing and append-reallocation surprises, struct embedding shadow bugs, generics constraint misuse, complexity, and temporal history — plus tree-sitter-based crash-class detectors (index-out-of-range on unchecked slice access, nil-map-write, integer overflow on untrusted input, `sync.WaitGroup` misuse), goroutine-leak detectors (fire-and-forget goroutines with no cancellation path), a differential parity checker against a project's own generated or hand-duplicated code (e.g. mocks vs interfaces drifting apart), and a stress reproducer that turns a static race/leak candidate into a confirmed failure under `-race` and load.

**8 commands** (`explore`, `informed-explore`, `known-issues`, `reproduce`, `map`, `hotspots`, `health`, `diff-review`) for different analysis workflows.

**Analysis scripts** — stdlib-only (`go/ast`, `go/parser`) scanners for the legacy dimensions, plus tree-sitter crash-class detectors, a known-issues regression checker, and an informed-explore briefing generator.

**Standalone CLI** (`goreview`) — every scanner and detector is also invocable outside Claude Code, for CI pipelines and pre-merge gates, emitting JSON, SARIF, or human-readable output.

## Reviewing the whole tree, one slice at a time

An `informed-explore` run works best when the whole scope fits in one context. As with the CPython toolkit this is modeled on, large packages strain a single dispatch and mid-size slices triage well — so for any repository above a size threshold, the reviewable surface is partitioned into bounded slices (by line count, not file count, since Go files vary wildly in density) and tracked in `data/review_slices.json`.

```bash
python tools/slice_status.py --catalog-dir ~/projects/go-review-findings
python tools/make_slice_context.py pkg-net --repo ~/projects/<target-repo> \
    --catalog-dir ~/projects/go-review-findings
/go-review-toolkit:informed-explore internal/net all
```

`slice_status.py` is the campaign cursor — progress per tier, findings per slice, and the next slice to run. `make_slice_context.py` does the setup that would otherwise be manual: the run tree, the informed briefing, every scanner run both corpus-wide and re-run slice-scoped (never post-filtered, so denominators stay honest), the calibration/new-territory split, and the `RUN_CONTEXT.md` the agents read. When a slice finishes, flip its status to done and commit.

A target repo gains and loses files as it evolves, so the manifest can silently stop covering the tree. `--verify` re-walks a checkout and fails on any unassigned or vanished file; `--sync` refreshes line counts.

## Prerequisites

- Claude Code installed and running.
- Go 1.21+ for the analysis scripts and standalone CLI.
- `tree-sitter` + `tree-sitter-go` (or the Go module equivalent) for the crash-class and race detectors, `known-issues`, and `informed-explore`; the legacy scanners remain stdlib-only (`go/ast`, `go/parser`, `go/types`).
- `go test -race` support in the target module, for the `reproduce` command.

## How It Works

Two generations of scanner coexist: stdlib-only `go/ast`-based scanners for the legacy dimensions (errors, context, defer, interfaces, complexity), and tree-sitter detectors that parse a real Go syntax tree to target specific reachable crash and race classes (validated against confirmed issues in real Go projects). All scripts report **candidates** — expect a meaningful false-positive rate depending on the detector — and the agents read the actual code to confirm or dismiss each finding and classify it FIX / CONSIDER / POLICY / ACCEPTABLE.

Beyond static analysis, the `reproduce` command closes the loop: it runs the target under `go test -race`, sustained load, and (where applicable) `GOMAXPROCS`/scheduler-pressure sweeps to turn a static candidate into a reproduced failure — with the discipline that a non-reproduction is reported honestly and does not by itself refute the finding.

For detailed usage, agent descriptions, and recommended workflows, see the plugin README.

## License

MIT — see LICENSE for details.# go-review-toolkit
