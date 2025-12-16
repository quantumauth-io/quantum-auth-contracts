// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import { ITPMVerifier } from "../interfaces/ITPMVerifier.sol";

contract MockTPMVerifier is ITPMVerifier {
    bytes32 public expectedKeyId;
    bytes32 public expectedHash;
    bytes public expectedSig;

    function setExpected(bytes32 keyId, bytes32 messageHash, bytes calldata sig) external {
        expectedKeyId = keyId;
        expectedHash = messageHash;
        expectedSig = sig;
    }

    function verify(
        bytes32 keyId,
        bytes32 messageHash,
        bytes calldata signature
    ) external view returns (bool ok) {
        return (keyId == expectedKeyId &&
        messageHash == expectedHash &&
            keccak256(signature) == keccak256(expectedSig));
    }
}
