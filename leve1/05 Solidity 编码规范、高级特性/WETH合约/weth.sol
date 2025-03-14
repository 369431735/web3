// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract WETH {
    string public name = "Wrapped Ether";
    string public symbol = "WETH";
    uint8 public decimals = 18;

    event Approval(address indexed src, address indexed delegateAds, uint256 amount);
    event Transfer(address indexed src, address indexed toAds, uint256 amount);
    event Deposit(address indexed toAds, uint256 amount);
    event Withdraw(address indexed src, uint256 amount);

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    //存钱
    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }
  //取钱
    function withdraw(uint256 amount) public {
        require(balanceOf[msg.sender] >= amount);
        balanceOf[msg.sender] -= amount;
        (bool success, ) = payable(msg.sender).call{value: amount}("");
        require(success, "Transfer failed");
        emit Withdraw(msg.sender, amount);
    }
//获取总钱数
    function totalSupply() public view returns (uint256) {
        return address(this).balance;
    }
//授权功能
    function approve(address delegateAds, uint256 amount) public returns (bool) {
        allowance[msg.sender][delegateAds] = amount;
        emit Approval(msg.sender, delegateAds, amount);
        return true;
    }
//转账
    function transfer(address toAds, uint256 amount) public returns (bool) {
        return transferFrom(msg.sender, toAds, amount);
    }

    function transferFrom(
        address src,
        address toAds,
        uint256 amount
    ) public returns (bool) {
        require(balanceOf[src] >= amount);
        if (src != msg.sender) {
            require(allowance[src][msg.sender] >= amount);
            allowance[src][msg.sender] -= amount;
        }
        balanceOf[src] -= amount;
        balanceOf[toAds] += amount;
        emit Transfer(src, toAds, amount);
        return true;
    }


}