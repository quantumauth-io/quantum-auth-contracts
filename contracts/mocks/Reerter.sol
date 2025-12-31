// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

contract Reverter {
    error Nope();

    function boom() external pure {
        revert Nope();
    }
}