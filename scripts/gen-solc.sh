#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

# Ensure forge exists
FORGE="$(command -v forge || true)"
if [ -z "$FORGE" ]; then
  echo "forge not found."
  echo "Install with: curl -L https://foundry.paradigm.xyz | bash && foundryup"
  exit 1
fi

# Build with Foundry (expects foundry.toml in repo root of contracts)
"$FORGE" --version
"$FORGE" build

mkdir -p abi bin

# Helper: copy ABI+BIN from Foundry out/ tree into our stable locations
copy_artifact () {
  local contract="$1"          # e.g. QuantumAuthAccount
  local type_name="$2"         # e.g. QuantumAuthAccount (same as contract unless you use different type)
  local out_dir="$3"           # e.g. out/QuantumAuthAccount.sol
  local json="$out_dir/$type_name.json"

  if [ ! -f "$json" ]; then
    echo "Missing Foundry artifact: $json"
    echo "Tip: check contract name/type and Foundry output structure."
    exit 1
  fi

  # ABI
  jq -c '.abi' "$json" > "abi/$contract.abi.json"

  # Bytecode object (standard-json output)
  jq -r '.bytecode.object' "$json" > "bin/$contract.bin"

  # Sanity: ensure non-empty
  test -s "abi/$contract.abi.json"
  test -s "bin/$contract.bin"
}

# You must align these paths with where the contracts actually live in your repo.
# Adjust the out_dir if your solidity files are in different folders.
copy_artifact "QuantumAuthAccount" "QuantumAuthAccount" "out/QuantumAuthAccount.sol"
copy_artifact "TPMVerifierSecp256k1" "TPMVerifierSecp256k1" "out/TPMVerifierSecp256k1.sol"
copy_artifact "QAERC20" "QAERC20" "out/QAERC20.sol"
copy_artifact "EntryPoint" "EntryPoint" "out/EntryPoint.sol"

echo "Artifacts:"
ls -la abi bin | sed -n '1,200p'
