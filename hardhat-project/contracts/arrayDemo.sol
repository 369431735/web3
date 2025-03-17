// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract arrayDemo {
    // 定义一个公共的字符串变量 message
   int[] public array;

    function put(int i) public {
       array.put(i);
    }

    // 获取 message 变量的值
    function getArray() public view returns (int[]) {
        return array;
    }
}