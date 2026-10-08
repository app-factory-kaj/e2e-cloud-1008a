# Greeting

## Purpose

Returns a JSON greeting for a caller-supplied name via `GET /hello`.

## User Stories

- F1.1 As an API caller, I send `GET /hello?name=<name>` and receive a 200 JSON response with a greeting for that name.
- F1.2 As an API caller, I send `GET /hello` without a `name` parameter and receive a 200 JSON response with a generic greeting ("Hello, World!") instead of an error.
- F1.3 As an API caller, I receive the greeting as a small JSON object with a single message field (e.g. `{"message": "Hello, Alice!"}`), so my client can parse it reliably. *assumed*

## Decisions

- `/hello` requires no authentication, login or sign-in of any kind; it is open to any caller.
- The service keeps no database or persistent storage; it holds no state between requests, in memory or otherwise.
- The service follows the conventions demonstrated in `app-factory-kaj/e2e-reference`.
- The `name` value is used as given, with no format validation beyond being a string. *assumed*

## Out of Scope

- No user interface — this feature is API-only.
- No persistence of names or greetings.
- No authentication or authorization.