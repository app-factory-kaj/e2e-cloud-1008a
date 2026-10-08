# Product-wide

Rules that apply to more than one feature.

This product has a single feature today, so there are no cross-feature
requirements yet. This file stays the entry point for any that emerge later.

## Decisions

- No authentication, login or sign-in of any kind (no Thunder/SSO) — every endpoint is open to any caller.
- No database or persistent storage of any kind. Any state the service holds is kept in memory only and is lost on restart.