// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Lottery {
    // 管理员地址
    address  public manager;
    // 玩家地址
    address[]  public players;

    // 构造函数，在合约部署时初始化 message 变量
    constructor(address memory _message) {
        message = _message;
    }

    // 设置 manager 变量的值
    function setManager(address memory _manager) public {
        manager = _manager;
    }

    // 获取 manager 变量的值
    function getManager() public view returns (address ) {
        return manager;
    }


    function random() private view returns (unit){
        return uint(keccak256(block.difficulty,now,players));
    }
    //
    function pickWinner() public  onlyMangerCanCallCall returns (address ) {
        uint index=random()%players.length;
        address  winner=players[index];
        winner.transfer(this.balance);
        return winner;
    }

    modifier onlyMangerCanCallCall(){
        require(msg.sender == manager);
        _;
    }
}