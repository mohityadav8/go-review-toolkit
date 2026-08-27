// Package campaign implements the review-slice campaign cursor: tracking
// which slices of a large repository have been reviewed, which are
// pending, and what findings each slice produced. Mirrors the role of
// tools/slice_status.py and tools/make_slice_context.py in the reference
// CPython toolkit, but as a Go package the CLI can call directly.
package campaign

// TODO: define Slice, Campaign, and Status types, and the functions that
// read/write data/review_slices/*.json.
