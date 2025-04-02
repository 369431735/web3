 // SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

/**
 * @notice 多重签名接口定义
 */
interface IMultiSignature {
    /**
     * @notice 获取有效签名
     * @param msghash 消息哈希
     * @param lastIndex 上次索引
     * @return 新的签名索引
     */
    function getValidSignature(bytes32 msghash, uint256 lastIndex) external view returns(uint256);
}

/**
 * @title 多重签名客户端合约
 * @notice 为其他合约提供多重签名功能的基础合约
 */
contract multiSignatureClient {
    // 多重签名存储位置，使用keccak256哈希确定唯一位置
    uint256 private constant multiSignaturePositon = uint256(keccak256("org.multiSignature.storage"));
    // 默认索引值
    uint256 private constant defaultIndex = 0;

    /**
     * @notice 构造函数，初始化多重签名合约地址
     * @param multiSignature 多重签名合约地址
     */
    constructor(address multiSignature) public {
        require(multiSignature != address(0), "multiSignatureClient : Multiple signature contract address is zero!");
        saveValue(multiSignaturePositon, uint256(multiSignature));
    }

    /**
     * @notice 获取多重签名合约地址
     * @return 多重签名合约地址
     */
    function getMultiSignatureAddress() public view returns (address) {
        return address(getValue(multiSignaturePositon));
    }

    /**
     * @notice 验证调用是否经过多重签名
     */
    modifier validCall() {
        checkMultiSignature();
        _;
    }

    /**
     * @notice 检查多重签名状态
     * @dev 使用内联汇编获取调用值，并验证交易是否已获得批准
     */
    function checkMultiSignature() internal view {
        uint256 value;
        assembly {
            value := callvalue()
        }
        // 生成消息哈希
        bytes32 msgHash = keccak256(abi.encodePacked(msg.sender, address(this)));
        address multiSign = getMultiSignatureAddress();
        // 获取新的签名索引
        uint256 newIndex = IMultiSignature(multiSign).getValidSignature(msgHash, defaultIndex);
        require(newIndex > defaultIndex, "multiSignatureClient : This tx is not aprroved");
    }

    /**
     * @notice 保存数据到特定存储位置
     * @param position 存储位置
     * @param value 要存储的值
     */
    function saveValue(uint256 position, uint256 value) internal {
        assembly {
            sstore(position, value)
        }
    }

    /**
     * @notice 从特定存储位置读取数据
     * @param position 存储位置
     * @return value 存储的值
     */
    function getValue(uint256 position) internal view returns (uint256 value) {
        assembly {
            value := sload(position)
        }
    }
}
