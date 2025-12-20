#!/bin/sh
set -eu

if ! command -v goose >/dev/null 2>&1; then
  echo "goose is required but not installed" >&2
  exit 1
fi

if [ -z "${POSTGRES_DSN:-}" ]; then
  echo "POSTGRES_DSN is required to run migrations" >&2
  exit 1
fi

MIGRATIONS_DIR="${MIGRATIONS_DIR:-$(cd "$(dirname "$0")/.." && pwd)/migrations}"

goose -dir "$MIGRATIONS_DIR" postgres "$POSTGRES_DSN" up
