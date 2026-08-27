---
name: task-workflow
description: Shared workflow guidance for every go-review-toolkit agent — how to read a slice, classify a finding, and write it to reports/.
---

# Task workflow

This skill is loaded by every agent in `plugin/agents/`. It defines the
shared review workflow so agents don't each reinvent it:

1. Read the assigned scope (a slice, a single file, or a diff).
2. Run the relevant scanner(s) and read their candidate output.
3. For each candidate, read the actual code and classify it as
   **FIX / CONSIDER / POLICY / ACCEPTABLE**.
4. Write findings to `reports/` in the shared format.

<!-- TODO: fill in the exact report schema and file-naming convention -->
