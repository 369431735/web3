// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract FunctionModifier {
    // We will use these variables to demonstrate how to use
    // modifiers.
    address public owner;
    uint256 public x = 10;
    bool public locked;
//构造器
    constructor() {
        owner = msg.sender;
    }


    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }


    modifier validAddress(address _addr) {
        require(_addr != address(0), "Not valid address");
        _;
    }

    function changeOwner(address _newOwner)
    public
    onlyOwner
    validAddress(_newOwner)
    {
        owner = _newOwner;
    }

   //加锁 解决并发
    modifier noReentrancy() {
        require(!locked, "No reentrancy");
        locked = true;
        _;
        locked = false;
    }

    function decrement(uint256 i) public noReentrancy {
        x -= i;

        if (i > 1) {
            decrement(i - 1);
        }
    }
}