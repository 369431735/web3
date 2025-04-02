// SPDX-License-Identifier: MIT
pragma solidity 0.6.12;

import "./multiSignatureClient.sol";

// 白名单地址操作库
library whiteListAddress {
    // 添加地址到白名单
    function addWhiteListAddress(address[] storage whiteList, address temp) internal {
        if (!isEligibleAddress(whiteList, temp)) {
            whiteList.push(temp);
        }
    }

    // 从白名单中移除地址
    function removeWhiteListAddress(address[] storage whiteList, address temp) internal returns (bool) {
        uint256 len = whiteList.length;
        uint256 i = 0;
        for (; i < len; i++) {
            if (whiteList[i] == temp)
                break;
        }
        if (i < len) {
            if (i != len-1) {
                whiteList[i] = whiteList[len-1];
            }
            whiteList.pop();
            return true;
        }
        return false;
    }

    // 检查地址是否在白名单中
    function isEligibleAddress(address[] memory whiteList, address temp) internal pure returns (bool) {
        uint256 len = whiteList.length;
        for (uint256 i = 0; i < len; i++) {
            if (whiteList[i] == temp)
                return true;
        }
        return false;
    }
}

// 多重签名合约
contract multiSignature is multiSignatureClient {
    uint256 private constant defaultIndex = 0;
    using whiteListAddress for address[];

    // 签名者列表
    address[] public signatureOwners;
    // 最小签名数阈值
    uint256 public threshold;

    // 签名信息结构
    struct signatureInfo {
        address applicant;      // 申请人
        address[] signatures;   // 已签名列表
    }

    // 消息哈希 => 签名信息数组
    mapping(bytes32 => signatureInfo[]) public signatureMap;

    // 事件声明
    event TransferOwner(address indexed sender, address indexed oldOwner, address indexed newOwner);
    event CreateApplication(address indexed from, address indexed to, bytes32 indexed msgHash);
    event SignApplication(address indexed from, bytes32 indexed msgHash, uint256 index);
    event RevokeApplication(address indexed from, bytes32 indexed msgHash, uint256 index);

    // 构造函数：初始化签名者列表和阈值
    constructor(address[] memory owners, uint256 limitedSignNum) multiSignatureClient(address(this)) public {
        require(owners.length >= limitedSignNum, "Multiple Signature : Signature threshold is greater than owners' length!");
        signatureOwners = owners;
        threshold = limitedSignNum;
    }

    // 转移签名者权限
    function transferOwner(uint256 index, address newOwner) public onlyOwner validCall {
        require(index < signatureOwners.length, "Multiple Signature : Owner index is overflow!");
        emit TransferOwner(msg.sender, signatureOwners[index], newOwner);
        signatureOwners[index] = newOwner;
    }

    // 创建签名申请
    function createApplication(address to) external returns(uint256) {
        bytes32 msghash = getApplicationHash(msg.sender, to);
        uint256 index = signatureMap[msghash].length;
        signatureMap[msghash].push(signatureInfo(msg.sender, new address[](0)));
        emit CreateApplication(msg.sender, to, msghash);
        return index;
    }

    // 签名申请
    function signApplication(bytes32 msghash) external onlyOwner validIndex(msghash, defaultIndex) {
        emit SignApplication(msg.sender, msghash, defaultIndex);
        signatureMap[msghash][defaultIndex].signatures.addWhiteListAddress(msg.sender);
    }

    // 撤销签名
    function revokeSignApplication(bytes32 msghash) external onlyOwner validIndex(msghash, defaultIndex) {
        emit RevokeApplication(msg.sender, msghash, defaultIndex);
        signatureMap[msghash][defaultIndex].signatures.removeWhiteListAddress(msg.sender);
    }

    // 获取有效签名
    function getValidSignature(bytes32 msghash, uint256 lastIndex) external view returns(uint256) {
        signatureInfo[] storage info = signatureMap[msghash];
        for (uint256 i = lastIndex; i < info.length; i++) {
            if(info[i].signatures.length >= threshold) {
                return i + 1;
            }
        }
        return 0;
    }

    // 获取申请信息
    function getApplicationInfo(bytes32 msghash, uint256 index) validIndex(msghash, index) public view returns (address, address[]memory) {
        signatureInfo memory info = signatureMap[msghash][index];
        return (info.applicant, info.signatures);
    }

    // 获取申请数量
    function getApplicationCount(bytes32 msghash) public view returns (uint256) {
        return signatureMap[msghash].length;
    }

    // 获取申请哈希
    function getApplicationHash(address from, address to) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(from, to));
    }

    // 仅限签名者调用
    modifier onlyOwner {
        require(signatureOwners.isEligibleAddress(msg.sender), "Multiple Signature : caller is not in the ownerList!");
        _;
    }

    // 验证索引有效性
    modifier validIndex(bytes32 msghash, uint256 index) {
        require(index < signatureMap[msghash].length, "Multiple Signature : Message index is overflow!");
        _;
    }
}