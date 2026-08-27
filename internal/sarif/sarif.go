// Package sarif implements SARIF (Static Analysis Results Interchange
// Format) output, so goreview's findings can be consumed by GitHub code
// scanning and other tools that understand the SARIF spec.
package sarif

// TODO: implement the SARIF 2.1.0 schema subset needed to represent a
// go-review-toolkit finding (rule, location, message, level).
