// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract arrayDemo {
    // 定义一个整数数组
    int[] public array;

    // 添加元素到数组
    function put(int i) public {
        array.push(i);
    }

    // 获取整个数组
    function getArray() public view returns (int[] memory) {
        return array;
    }

    // 获取数组长度
    function getLength() public view returns (uint) {
        return array.length;
    }

    // 根据索引获取元素
    function getElement(uint index) public view returns (int) {
        require(index < array.length, "Index out of bounds");
        return array[index];
    }
}