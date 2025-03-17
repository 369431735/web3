pragma solidity ^0.8.0;

contract Purchase {
    uint public value;
    address public seller;
    address public buyer;
    enum State { Created, Locked, Inactive }
    State public state;

    constructor() public payable {
        seller = msg.sender;
        value = msg.value / 2;
        require((2 * value) == msg.value, "Value has to be even.");
    }

    modifier condition(bool _condition) {
        require(_condition);
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

    event Aborted();
    event PurchaseConfirmed();
    event ItemReceived();

    function abort()
    public
    onlySeller
    inState(State.Created)
    {
        emit Aborted();
        state = State.Inactive;
        (bool success, ) = seller.call{value: address(this).balance}("");
        require(success, "Transfer failed");
    }

    function confirmPurchase()
    public
    inState(State.Created)
    condition(msg.value == (2 * value))
    payable
    {
        emit PurchaseConfirmed();
        buyer = msg.sender;
        state = State.Locked;
    }

    function confirmReceived() public onlyBuyer inState(State.Locked)
    {
        emit ItemReceived();
        state = State.Inactive; // 先更新状态
        (bool success1, ) = buyer.call{value: value}("");
        require(success1, "Transfer to buyer failed");
        (bool success2, ) = seller.call{value: address(this).balance}("");
        require(success2, "Transfer to seller failed");
    }
}