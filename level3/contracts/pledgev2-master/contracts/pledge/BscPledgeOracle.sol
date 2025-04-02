// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "../multiSignature/multiSignatureClient.sol";
import "./AggregatorV3Interface.sol";

/**
 * @title BSC抵押预言机合约
 * @dev 用于获取和管理资产价格的预言机合约
 */
contract BscPledgeOracle is multiSignatureClient {
    // 价格聚合器接口映射：资产地址 => 预言机接口
    mapping(uint256 => AggregatorV3Interface) public assetsMap;
    // 精度映射：资产地址 => 精度
    mapping(uint256 => uint8) public decimalsMap;

    /**
     * @notice 构造函数
     * @param multiSignature 多重签名合约地址
     */
    constructor(address multiSignature) multiSignatureClient(multiSignature) public {
    }

    /**
     * @notice 批量设置价格聚合器
     * @dev 只能通过多重签名调用
     * @param _assets 资产地址数组
     * @param _aggregators 价格聚合器地址数组
     */
    function setAssets(uint256[] memory _assets, address[] memory _aggregators) external validCall {
        require(_assets.length == _aggregators.length, "BscPledgeOracle: 数组长度不匹配");
        for(uint i = 0; i < _assets.length; i++) {
            assetsMap[_assets[i]] = AggregatorV3Interface(_aggregators[i]);
        }
    }

    /**
     * @notice 批量设置资产精度
     * @dev 只能通过多重签名调用
     * @param _assets 资产地址数组
     * @param _decimals 精度数组
     */
    function setDecimals(uint256[] memory _assets, uint8[] memory _decimals) external validCall {
        require(_assets.length == _decimals.length, "BscPledgeOracle: 数组长度不匹配");
        for(uint i = 0; i < _assets.length; i++) {
            decimalsMap[_assets[i]] = _decimals[i];
        }
    }

    /**
     * @notice 批量获取资产价格
     * @param _assets 资产地址数组
     * @return prices 价格数组
     */
    function getPrices(uint256[] memory _assets) external view returns(uint256[] memory prices) {
        prices = new uint256[](_assets.length);
        for(uint i = 0; i < _assets.length; i++) {
            // 获取单个资产价格
            (, int256 price,,,) = assetsMap[_assets[i]].latestRoundData();
            // 如果价格小于0，认为价格无效
            require(price > 0, "BscPledgeOracle: 无效价格");
            // 转换价格格式并存储
            prices[i] = uint256(price);
        }
    }
}
