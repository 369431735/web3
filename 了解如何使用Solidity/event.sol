// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract Event {
    // 申明事件
    event Log(address indexed sender, string message);
    event AnotherLog();

    function test() public {
        //触发事件
        emit Log(msg.sender, "Hello World!");
        emit Log(msg.sender, "Hello EVM!");
        emit AnotherLog();
    }
}