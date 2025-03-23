// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ERC20MinerReward is ERC20 {
    // Added 'event' keyword and proper syntax for event declaration
    event LogNewAlert(
        string description,
        address indexed _from,
        uint256 _n
    );

    constructor() ERC20("MinerReward", "MRW") {}

    function reward() public {
        _mint(block.coinbase, 20);
        // Changed single quotes to double quotes for string literal
        emit LogNewAlert("_rewarded", block.coinbase, block.number);
    }
}