// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract ViewAndPure  {

    uint256 public x = 1;

    // view 只读不修改
    function addToX(uint256 y) public view returns (uint256) {
        return x + y;
    }

    // pure 不读不修改
    function add(uint256 i, uint256 j) public pure returns (uint256) {
        return i + j;
    }

}
