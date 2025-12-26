#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

# Ensure abigen exists
command -v abigen >/dev/null 2>&1 || {
  echo "abigen not found. Install with: go install github.com/ethereum/go-ethereum/cmd/abigen@v1.16.7"
  exit 1
}

# Sanity: ensure artifacts exist
test -f abi/QuantumAuthAccount.abi.json
test -f bin/QuantumAuthAccount.bin
test -f abi/TPMVerifierSecp256k1.abi.json
test -f bin/TPMVerifierSecp256k1.bin

mkdir -p bindings/go/quantumauthaccount
mkdir -p bindings/go/tpmverifiersecp256k1

abigen \
  --abi abi/QuantumAuthAccount.abi.json \
  --bin bin/QuantumAuthAccount.bin \
  --pkg quantumauthaccount \
  --type QuantumAuthAccount \
  --out bindings/go/quantumauthaccount/quantumauthaccount.go

abigen \
  --abi abi/TPMVerifierSecp256k1.abi.json \
  --bin bin/TPMVerifierSecp256k1.bin \
  --pkg tpmverifiersecp256k1 \
  --type TPMVerifierSecp256k1 \
  --out bindings/go/tpmverifiersecp256k1/tpmverifiersecp256k1.go
