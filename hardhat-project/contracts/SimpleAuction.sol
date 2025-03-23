// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract SimpleAuction {
    address public beneficiary;
    uint256 public auctionEndTime;
    address public highestBidder;
    uint256 public highestBid;
    mapping(address => uint256) public pendingReturns;
    bool public ended;

    event HighestBidIncreased(address bidder, uint256 amount);
    event AuctionEnded(address winner, uint256 amount);

    constructor(uint256 _biddingTime, address _beneficiary) {
        require(_beneficiary != address(0), "Beneficiary cannot be zero address");
        beneficiary = _beneficiary;
        auctionEndTime = block.timestamp + _biddingTime;
    }

    function bid() public payable {
        require(!ended, "Auction already ended");
        require(block.timestamp < auctionEndTime, "Bid after auction end");
        require(msg.value > highestBid, "Insufficient bid");

        if (highestBidder != address(0)) {
            pendingReturns[highestBidder] += highestBid;
        }

        highestBidder = msg.sender;
        highestBid = msg.value;
        emit HighestBidIncreased(msg.sender, msg.value);
    }

    function withdraw() public {
        uint256 amount = pendingReturns[msg.sender];
        if (amount == 0) return;

        // 使用 send 代替 call（自动检查返回值）
        bool success = payable(msg.sender).send(amount);
        require(success, "Withdrawal failed");

        pendingReturns[msg.sender] = 0;
    }

    function auctionEnd() public {
        require(ended, "Auction not ended");
        require(block.timestamp >= auctionEndTime, "Auction not finished");

        ended = true;
        address winner = highestBidder;
        uint256 prize = highestBid;

        // 使用 transfer 代替 call（自动抛出异常）
        payable(beneficiary).transfer(prize);

        emit AuctionEnded(winner, prize);
    }
}