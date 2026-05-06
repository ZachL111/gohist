# Failure Modes

For `gohist`, I would look first for these mistakes:

- `span volume` cases moving lanes without a matching threshold change.
- `signal loss` scoring higher after drag increases.
- Duplicate fixture ids hiding a stale golden row.
- README examples drifting away from the verifier.

The local checks are intentionally strict about these cases.
