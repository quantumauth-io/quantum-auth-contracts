#!/usr/bin/env bash
set -euo pipefail

SOLC_IMAGE="${SOLC_IMAGE:-ghcr.io/argotorg/solc:0.8.28}"

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

mkdir -p build/solc abi bin

# Common args
ARGS=(
  --base-path .
  --include-path node_modules
  --include-path contracts
  --optimize --optimize-runs 200
  --abi --bin
  --overwrite
  -o build/solc
  contracts/account/QuantumAuthAccount.sol
  contracts/TPMVerifierSecp256k1.sol
)

# Try as ENTRYPOINT (image runs solc by default)
if docker run --rm \
  -u "$(id -u)":"$(id -g)" \
  -v "$ROOT:/workspace" \
  -w /workspace \
  "$SOLC_IMAGE" \
  "${ARGS[@]}"; then
  echo "solc via image entrypoint OK"
else
  echo "entrypoint mode failed; trying explicit 'solc' command..."
  docker run --rm \
    -u "$(id -u)":"$(id -g)" \
    -v "$ROOT:/workspace" \
    -w /workspace \
    "$SOLC_IMAGE" \
    solc "${ARGS[@]}"
fi

# Normalize output locations
cp -f build/solc/QuantumAuthAccount.abi abi/QuantumAuthAccount.abi.json
cp -f build/solc/QuantumAuthAccount.bin bin/QuantumAuthAccount.bin

cp -f build/solc/TPMVerifierSecp256k1.abi abi/TPMVerifierSecp256k1.abi.json
cp -f build/solc/TPMVerifierSecp256k1.bin bin/TPMVerifierSecp256k1.bin

echo "Artifacts:"
ls -la build/solc abi bin | sed -n '1,200p'
