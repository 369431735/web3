// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract PrimitiveDataTypes{
    /**
    * 无符号数  min=0    uint8 :max =mod(2,7)-1 uint256 :max =mod(2,255)-1  uint equals  uint256
    **/
    uint8 public u8 = 1;
    uint256 public u256 = 456;
    uint public u = 123;
/**
*二进制有符号数 首位0 正数
*/
    int8 public w8 = -1;
    int256 public w256 = 456;
    int public w = -123;
//各种类型取最大值 最小值
    int256 public minInt = type(int256).min;
    int256 public maxInt = type(int256).max;
    int8 public minInt = type(int8).min;
    int8 public maxInt = type(int8).max;

    //账户地址
    address public addr = 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c;
   //0x开头16进制
    bytes1 a = 0xb5; //  [10110101]
    bytes1 b = 0x56; //  [01010110]

    // Default values
    // Unassigned variables have a default value
    bool public defaultBoo; // false
    uint256 public defaultUint; // 0
    int256 public defaultInt; // 0
    address public defaultAddr; // 0x0000000000000000000000000000000000000000

}
