// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract EtherWallet {
    address payable public immutable owner;

    event Log(string funName, address from, uint256 value, bytes data);
    event Withdrawn(address indexed owner, uint256 amount);

    constructor() {
        owner = payable(msg.sender);
    }

    receive() external payable {
        emit Log("receive", msg.sender, msg.value, "");
    }

    // 提取固定金额（使用 transfer）
    function withdrawTransfer() external {
        require(msg.sender == owner, "Not owner");
        payable(msg.sender).transfer(100);
        emit Withdrawn(msg.sender, 100);
    }

    // 提取固定金额（使用 send）
    function withdrawSend() external {
        require(msg.sender == owner, "Not owner");
        bool success = payable(msg.sender).send(200);
        require(success, "Send Failed");
        emit Withdrawn(msg.sender, 200);
    }

    // 提取全部余额（使用 call）
    function withdrawCall() external {
        require(msg.sender == owner, "Not owner");
        uint256 balance = address(this).balance;
        (bool success, ) = msg.sender.call{value: balance}("");
        require(success, "Call Failed");
        emit Withdrawn(msg.sender, balance);
    }

    // 提取任意金额
    function withdraw(uint256 amount) external {
        require(msg.sender == owner, "Not owner");
        require(address(this).balance >= amount, "Insufficient balance");
        (bool success, ) = msg.sender.call{value: amount}("");
        require(success, "Call Failed");
        emit Withdrawn(msg.sender, amount);
    }

    // 查询余额
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}