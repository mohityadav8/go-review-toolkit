# Reproducer

Turns a static race/leak candidate into a confirmed failure by running the
target under `go test -race`, sustained load, and (where applicable)
`GOMAXPROCS`/scheduler-pressure sweeps.

A non-reproduction is reported honestly and does not by itself refute the
underlying static finding.

<!-- TODO: implement the stress-run harness and result reporting -->
