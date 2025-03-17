// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Inbox {
    // 定义一个公共的字符串变量 message
    string public message;

    // 构造函数，在合约部署时初始化 message 变量
    constructor(string memory _message) {
        message = _message;
    }

    // 设置 message 变量的值
    function setMessage(string memory _message) public {
        message = _message;
    }

    // 获取 message 变量的值
    function getMessage() public view returns (string memory) {
        return message;
    }
}