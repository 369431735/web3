// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract Immutable{
    /**
   * 不可变变量 部署时赋值 不可更改 读取时 Gas 消耗较低
   **/
    uint256 public immutable MAX_SUPPLY_IM;
    address public immutable OWNER;
    uint256 public immutable CREATION_TIMESTAMP;

    constructor(uint256 _MAX_SUPPLY_IM) {
        MAX_SUPPLY_IM = _MAX_SUPPLY_IM;
        OWNER = msg.sender;
        CREATION_TIMESTAMP = block.timestamp;
    }

}
