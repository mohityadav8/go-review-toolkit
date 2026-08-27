# Tools

Campaign and setup scripts, mirroring `tools/slice_status.py` and
`tools/make_slice_context.py` from the reference CPython toolkit:

- `slice_status` — the campaign cursor: progress per tier, findings per
  slice, and the next slice to run.
- `make_slice_context` — generates the run tree, the informed briefing,
  scanner runs (corpus-wide and slice-scoped), and `RUN_CONTEXT.md`.

<!-- TODO: implement both, in Go or Python to match your workflow -->
