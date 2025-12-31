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
  --include-path vendor
  --include-path vendor/account-abstraction
  --include-path vendor/account-abstraction/contracts
  --optimize --optimize-runs 200
  --abi --bin
  --overwrite
  -o build/solc
  contracts/account/QuantumAuthAccount.sol
  contracts/TPMVerifierSecp256k1.sol
  vendor/account-abstraction/contracts/core/EntryPoint.sol
  contracts/QAERC20.sol

)

# Try as ENTRYPOINT (image runs solc by default)
docker run --rm \
  -u "$(id -u)":"$(id -g)" \
  -v "$ROOT:/workspace" \
  -w /workspace \
  "$SOLC_IMAGE" \
  "${ARGS[@]}"


# Normalize output locations
cp -f build/solc/QuantumAuthAccount.abi abi/QuantumAuthAccount.abi.json
cp -f build/solc/QuantumAuthAccount.bin bin/QuantumAuthAccount.bin

cp -f build/solc/TPMVerifierSecp256k1.abi abi/TPMVerifierSecp256k1.abi.json
cp -f build/solc/TPMVerifierSecp256k1.bin bin/TPMVerifierSecp256k1.bin

cp -f build/solc/EntryPoint.abi abi/EntryPoint.abi.json
cp -f build/solc/EntryPoint.bin bin/EntryPoint.bin

cp -f build/solc/QAERC20.abi abi/QAERC20.abi.json
cp -f build/solc/QAERC20.bin bin/QAERC20.bin

echo "Artifacts:"
ls -la build/solc abi bin | sed -n '1,200p'
