// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract firstApp{
    uint256  public count;

    function add() public {
      count++;
    }
    function inc() public {
        count--;
    }

    function get() public view returns(uint256) {
        return count;
    }
}
