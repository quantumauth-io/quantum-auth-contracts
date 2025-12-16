#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

mkdir -p bindings/go/quantumauthaccount

# Ensure abigen exists
command -v abigen >/dev/null 2>&1 || {
  echo "abigen not found. Install with: go install github.com/ethereum/go-ethereum/cmd/abigen@v1.16.7"
  exit 1
}

abigen \
  --abi abi/QuantumAuthAccount.abi.json \
  --bin bin/QuantumAuthAccount.bin \
  --pkg quantumauthaccount \
  --type QuantumAuthAccount \
  --out bindings/go/quantumauthaccount/quantumauthaccount.go
