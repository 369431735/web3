// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract Purchase {
    uint public value;
    address public seller;
    address public buyer;
    enum State { Created, Locked, Inactive }
    State public state;

    event Aborted(address seller);
    event PurchaseConfirmed(address buyer, uint value);
    event ItemReceived(address buyer, address seller, uint value);

    constructor() payable {
        seller = msg.sender;
        value = msg.value / 2;
        require((2 * value) == msg.value, "Value has to be even.");
        state = State.Created;
    }

    modifier condition(bool _condition) {
        require(_condition, "Condition not met");
        _;
    }

    modifier onlyBuyer() {
        require(msg.sender == buyer, "Only buyer can call this.");
        _;
    }

    modifier onlySeller() {
        require(msg.sender == seller, "Only seller can call this.");
        _;
    }

    modifier inState(State _state) {
        require(state == _state, "Invalid state.");
        _;
    }

    function abort()
        public
        onlySeller
        inState(State.Created)
    {
        state = State.Inactive;
        payable(seller).transfer(address(this).balance);
        emit Aborted(seller);
    }

    function confirmPurchase()
        public
        payable
        inState(State.Created)
        condition(msg.value == (2 * value))
    {
        buyer = msg.sender;
        state = State.Locked;
        emit PurchaseConfirmed(buyer, msg.value);
    }

    function confirmReceived()
        public
        onlyBuyer
        inState(State.Locked)
    {
        state = State.Inactive;
        
        uint buyerRefund = value;
        uint sellerPayment = (2 * value) - buyerRefund;
        
        payable(buyer).transfer(buyerRefund);
        payable(seller).transfer(sellerPayment);

        emit ItemReceived(buyer, seller, value);
    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}