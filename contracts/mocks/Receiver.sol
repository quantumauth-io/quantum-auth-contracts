// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

contract Receiver {
    event Ping(uint256 x);
    uint256 public last;

    function ping(uint256 x) external {
        last = x;
        emit Ping(x);
    }
}
