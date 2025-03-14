// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Bank {
    // 状态变量
    address public immutable owner;

    // 事件
    event Deposit(address indexed from, uint256 amount);
    event Withdraw(uint256 amount);

    // 构造函数
    constructor() {
        owner = msg.sender;
    }

    // 接收以太币
    receive() external payable {
        require(msg.value > 0, "Deposit amount must be greater than 0");
        emit Deposit(msg.sender, msg.value);
    }

    // 提款
    function withdraw() external {
        require(msg.sender == owner, "Not Owner");
        uint256 balance = address(this).balance;
        emit Withdraw(balance);
        payable(msg.sender).transfer(balance);
    }

    // 获取余额
    function getBalance() external pure returns (uint256) {
        return address(this).balance;
    }

    // 销毁合约并发送余额
    function destroy() external {
        require(msg.sender == owner, "Only owner can destroy"); // 只有所有者可以销毁合约
        selfdestruct(payable(msg.sender)); // 销毁合约并发送余额
    }
}