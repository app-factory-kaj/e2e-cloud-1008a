# Greeter

## Problem Statement

Teams building services on this platform need a small, known-good reference service to confirm that new scaffolding follows the organization's conventions end-to-end. Without one, each new project has to re-derive basic HTTP service structure and conventions from scratch.

## Solution

Greeter is a small Go HTTP service that accepts a name and returns a JSON greeting, built to the conventions demonstrated in app-factory-kaj/e2e-reference. It serves as a minimal, working example of those conventions that other services can be checked against.

## Actors

- API Caller — any client that sends HTTP requests to the greeter service. No sign-in or identity is involved.

## Features

- F1 [Greeting](features/F1-greeting.md)

## Product-wide

See [Product-wide](product-wide.md).

## Out of Scope

- No user interface — this is an API-only service.
- No persistence or stored data.
- No authentication or authorization.