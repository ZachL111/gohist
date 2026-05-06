# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its observability focus without claiming live deployment or external usage.

## Cases

- `baseline`: `span volume`, score 135, lane `watch`
- `stress`: `latency skew`, score 188, lane `ship`
- `edge`: `signal loss`, score 196, lane `ship`
- `recovery`: `incident shape`, score 251, lane `ship`
- `stale`: `span volume`, score 197, lane `ship`

## Note

The repository should be understandable without pretending it is larger than it is.
