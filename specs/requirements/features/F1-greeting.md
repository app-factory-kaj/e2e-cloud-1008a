# Greeting

## Purpose

Returns a JSON greeting for a caller-supplied name via `GET /hello`.

## Decisions

- `/hello` requires no authentication; it is open to any caller.
- When the `name` query parameter is omitted, the service returns a generic greeting (e.g. "Hello, World!") rather than an error.