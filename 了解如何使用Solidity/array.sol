// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract Mapping {

    uint256[] public arr;  //空数组
    uint256[] public arr2 = [1, 2, 3];
    uint256[10] public myFixedSizeArr;  //固定长度数组

    //通过索引位置获取数组中的值
    function get(uint256 i) public view returns (uint256) {
        return arr[i];
    }
    //获取整个数组
    function getArr() public view returns (uint256[] memory) {
        return arr;
    }
    //往数组里面加值
    function push(uint256 i) public {
        arr.push(i);
    }
    //删除数组 arr 的最后一个元素，数组长度减少 1。
    function pop() public {
        arr.pop();
    }
    //获取数组长度
    function getLength() public view returns (uint256) {
        return arr.length;
    }
   //把索引index替换为默认值
    function remove(uint256 index) public {
        delete arr[index];
    }
    //在内存中创建一个固定大小为 5 的数组 a。
    //内存中的数组不能动态调整大小。
    function examples() external {
        uint256[] memory a = new uint256[](5);
    }
}
