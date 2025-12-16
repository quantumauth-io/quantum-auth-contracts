// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import { BaseAccount } from "@account-abstraction/contracts/core/BaseAccount.sol";
import { IEntryPoint } from "@account-abstraction/contracts/interfaces/IEntryPoint.sol";
import { PackedUserOperation } from "@account-abstraction/contracts/interfaces/PackedUserOperation.sol";

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { MessageHashUtils } from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import { ITPMVerifier } from "../interfaces/ITPMVerifier.sol";

/// @notice Immutable ERC-4337 account (PackedUserOperation) with 2-of-3 policy:
/// - Normal mode (mode=0): TPM + (EOA1 OR EOA2)
/// - Recovery mode (mode=1): EOA1 + EOA2 (TPM not required)
///
/// Security goal: No single compromised EOA can execute. No direct execute bypass.
contract QuantumAuthAccount is BaseAccount {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    // ---- signature modes ----
    uint8 internal constant MODE_NORMAL = 0;   // TPM + (EOA1 OR EOA2)
    uint8 internal constant MODE_RECOVERY = 1; // EOA1 + EOA2

    IEntryPoint private immutable _entryPoint;

    address public immutable eoa1;
    address public immutable eoa2;

    ITPMVerifier public immutable tpmVerifier;
    bytes32 public immutable tpmKeyId;

    error InvalidEOA();
    error InvalidTPMVerifier();
    error OnlyEntryPoint();
    error NotAuthorized();
    error InvalidMode();

    modifier onlyEntryPoint() {
        if (msg.sender != address(_entryPoint)) revert OnlyEntryPoint();
        _;
    }

    constructor(
        IEntryPoint entryPoint_,
        address eoa1_,
        address eoa2_,
        ITPMVerifier tpmVerifier_,
        bytes32 tpmKeyId_
    ) {
        if (eoa1_ == address(0) || eoa2_ == address(0) || eoa1_ == eoa2_) revert InvalidEOA();
        if (address(tpmVerifier_) == address(0)) revert InvalidTPMVerifier();

        _entryPoint = entryPoint_;
        eoa1 = eoa1_;
        eoa2 = eoa2_;

        tpmVerifier = tpmVerifier_;
        tpmKeyId = tpmKeyId_;
    }

    function entryPoint() public view override returns (IEntryPoint) {
        return _entryPoint;
    }

    /// @notice Execute a call from the account.
    /// @dev EntryPoint-only: prevents bypass of signature policy.
    function execute(address target, uint256 value, bytes calldata data) external override onlyEntryPoint {
        (bool ok, bytes memory ret) = target.call{ value: value }(data);
        if (!ok) {
            assembly {
                revert(add(ret, 0x20), mload(ret))
            }
        }
    }

    /// @notice Optional: withdraw ETH accidentally sent to the account.
    /// @dev Keep ownerless withdrawals out of MVP; require BOTH EOAs (recovery-style) via AA flow if needed later.
    function withdraw(address payable /*to*/, uint256 /*amount*/) external pure {
        revert NotAuthorized();
    }

    receive() external payable {}

    // ---------------- ERC-4337 validation ----------------

    /// @dev Enforces 2-of-3 policy using userOp.signature encoding:
    /// abi.encode(uint8 mode, bytes eoaSig1, bytes eoaSig2, bytes tpmSig)
    function _validateSignature(
        PackedUserOperation calldata userOp,
        bytes32 userOpHash
    ) internal view override returns (uint256 validationData) {
        (uint8 mode, bytes memory sig1, bytes memory sig2, bytes memory tpmSig) =
                            abi.decode(userOp.signature, (uint8, bytes, bytes, bytes));

        bytes32 ethHash = userOpHash.toEthSignedMessageHash();

        bool okEoa1 = _isValidEOASig(ethHash, sig1, eoa1);
        bool okEoa2 = _isValidEOASig(ethHash, sig2, eoa2);

        if (mode == MODE_NORMAL) {
            // TPM + (EOA1 OR EOA2)
            if (!(okEoa1 || okEoa2)) return 1;
            if (!tpmVerifier.verify(tpmKeyId, userOpHash, tpmSig)) return 1;
            return 0;
        }

        if (mode == MODE_RECOVERY) {
            // EOA1 + EOA2
            if (!(okEoa1 && okEoa2)) return 1;
            return 0;
        }

        revert InvalidMode();
    }

    function _isValidEOASig(bytes32 ethHash, bytes memory sig, address expected) internal pure returns (bool) {
        if (sig.length == 0) return false;
        return ethHash.recover(sig) == expected;
    }
}
