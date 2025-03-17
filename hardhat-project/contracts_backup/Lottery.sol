// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Lottery {
    // 管理员地址
    address public manager;
    // 玩家地址
    address[] public players;
    // 最小参与金额
    uint public constant MINIMUM_ENTRY_FEE = 0.01 ether;
    // 最近一次开奖时间
    uint public lastDrawTime;
    // 开奖冷却时间（1小时）
    uint public constant DRAW_COOLDOWN = 1 hours;

    // 事件
    event PlayerEntered(address player, uint amount);
    event WinnerPicked(address winner, uint amount);
    event LotteryReset(uint timestamp);

    // 构造函数，在合约部署时初始化管理员地址
    constructor() {
        manager = msg.sender;
        lastDrawTime = block.timestamp;
    }

    // 参与抽奖
    function enter() public payable {
        require(msg.value >= MINIMUM_ENTRY_FEE, "Minimum entry fee is 0.01 ether");
        players.push(msg.sender);
        emit PlayerEntered(msg.sender, msg.value);
    }

    // 生成随机数
    function random() private view returns (uint) {
        bytes32 hash = keccak256(
            abi.encodePacked(
                block.prevrandao,
                block.timestamp,
                block.number,
                players,
                blockhash(block.number - 1)
            )
        );
        return uint(hash);
    }

    // 选择获胜者
    function pickWinner() public onlyManagerCanCall returns (address) {
        require(players.length > 0, "No players in the lottery");
        require(block.timestamp >= lastDrawTime + DRAW_COOLDOWN, "Please wait for the cooldown period");

        uint index = random() % players.length;
        address winner = players[index];
        uint prizeAmount = address(this).balance;

        // 转账给获胜者
        payable(winner).transfer(prizeAmount);
        
        // 发出事件
        emit WinnerPicked(winner, prizeAmount);
        
        // 重置抽奖
        resetLottery();
        
        return winner;
    }

    // 重置抽奖
    function resetLottery() private {
        players = new address[](0);
        lastDrawTime = block.timestamp;
        emit LotteryReset(block.timestamp);
    }

    // 获取所有玩家
    function getPlayers() public view returns (address[] memory) {
        return players;
    }

    // 获取玩家数量
    function getPlayersCount() public view returns (uint) {
        return players.length;
    }

    // 获取奖池金额
    function getBalance() public view returns (uint) {
        return address(this).balance;
    }

    // 获取距离下次可开奖的剩余时间
    function getTimeUntilNextDraw() public view returns (uint) {
        uint nextDrawTime = lastDrawTime + DRAW_COOLDOWN;
        if (block.timestamp >= nextDrawTime) {
            return 0;
        }
        return nextDrawTime - block.timestamp;
    }

    // 紧急情况：只有在合约出现严重问题时使用
    function emergencyWithdraw() public onlyManagerCanCall {
        require(address(this).balance > 0, "No funds to withdraw");
        require(block.timestamp >= lastDrawTime + DRAW_COOLDOWN * 24, "Must wait 24 cooldown periods");
        payable(manager).transfer(address(this).balance);
    }

    // 只有管理员可以调用的修饰器
    modifier onlyManagerCanCall() {
        require(msg.sender == manager, "Only manager can call this function");
        _;
    }
}