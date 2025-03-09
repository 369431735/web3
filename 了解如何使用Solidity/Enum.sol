// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract Enum {

    // 状态枚举值
    enum Status {
        Pending,
        Shipped,
        Accepted,
        Rejected,
        Canceled
    }

    //status 是一个 Status 类型的公共变量。
    //未初始化时，status 的默认值是枚举的第一个值，即 Pending（0）。
    Status public status;

    // 获取当前状态
    // Pending  - 0
    // Shipped  - 1
    // Accepted - 2
    // Rejected - 3
    // Canceled - 4
    function get() public view returns (Status) {
        return status;
    }

    // 更新状态
    function set(Status _status) public {
        status = _status;
    }

    // 更新到特定状态
    function cancel() public {
        status = Status.Canceled;
    }

    // 重置状态 恢复到默认值Pending
    function reset() public {
        delete status;
    }

}
