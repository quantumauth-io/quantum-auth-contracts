#!/usr/bin/env bash
set -euo pipefail

SOLC_IMAGE="${SOLC_IMAGE:-ghcr.io/argotorg/solc:0.8.28}"

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

mkdir -p build/solc abi bin

# Compile ONLY the contracts you want bindings for (MVP: QuantumAuthAccount)
# include-path node_modules makes imports like @openzeppelin/... resolve to node_modules/@openzeppelin/...
docker run --rm \
  -v "$ROOT:/workspace" \
  -w /workspace \
  "$SOLC_IMAGE" \
  --base-path . \
  --include-path node_modules \
  --include-path contracts \
  --optimize --optimize-runs 200 \
  --abi --bin \
  --overwrite \
  -o build/solc \
  contracts/account/QuantumAuthAccount.sol

# Normalize output locations
# solc names files by contract name (e.g. QuantumAuthAccount.abi / QuantumAuthAccount.bin)
cp -f build/solc/QuantumAuthAccount.abi abi/QuantumAuthAccount.abi.json
cp -f build/solc/QuantumAuthAccount.bin bin/QuantumAuthAccount.bin
