// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract ArrayRemoveByShifting  {
    uint256[] public arr;

    //通过将索引i+1以后的数据复制到i的位置 删除最后一个元素
    function remove(uint256 _index) public {
        require(_index < arr.length, "index out of bound");
        for (uint256 i = _index; i < arr.length - 1; i++) {
            arr[i] = arr[i + 1];
        }
        arr.pop();
    }
  //将最后一个元素 替换到 i的位置
    function remove2(uint256 index) public {
        arr[index] = arr[arr.length - 1];
        arr.pop();
    }
}
