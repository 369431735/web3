// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract Constants{

 /**
 * 常量 编译时赋值  不可更改 读取时 Gas 消耗极低
 **/
    // 常量变量
    uint256 public constant MAX_SUPPLY = 1000000;
    address public constant OWNER = 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c;
    string public constant TOKEN_NAME = "MyToken";

}
