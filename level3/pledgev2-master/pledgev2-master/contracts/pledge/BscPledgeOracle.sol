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
}// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "../multiSignature/multiSignatureClient.sol";
import "@chainlink/contracts/src/v0.6/interfaces/AggregatorV3Interface.sol";



contract BscPledgeOracle is multiSignatureClient {

    mapping(uint256 => AggregatorV3Interface) internal assetsMap;
    mapping(uint256 => uint256) internal decimalsMap;
    mapping(uint256 => uint256) internal priceMap;
    uint256 internal decimals = 1;

    constructor(address multiSignature) multiSignatureClient(multiSignature) public {
//        //  bnb/USD
//        assetsMap[uint256(0x0000000000000000000000000000000000000000)] = AggregatorV3Interface(0x2514895c72f50D8bd4B4F9b1110F0D6bD2c97526);
//        // DAI/USD
//        assetsMap[uint256(0xf2bDB4ba16b7862A1bf0BE03CD5eE25147d7F096)] = AggregatorV3Interface(0xE4eE17114774713d2De0eC0f035d4F7665fc025D);
//        // BTC/USD
//        assetsMap[uint256(0xF592aa48875a5FDE73Ba64B527477849C73787ad)] = AggregatorV3Interface(0x5741306c21795FdCBb9b265Ea0255F499DFe515C);
//        // BUSD/USD
//        assetsMap[uint256(0xDc6dF65b2fA0322394a8af628Ad25Be7D7F413c2)] = AggregatorV3Interface(0x9331b55D9830EF609A2aBCfAc0FBCE050A52fdEa);
//
//
//        decimalsMap[uint256(0x0000000000000000000000000000000000000000)] = 18;
//        decimalsMap[uint256(0xf2bDB4ba16b7862A1bf0BE03CD5eE25147d7F096)] = 18;
//        decimalsMap[uint256(0xF592aa48875a5FDE73Ba64B527477849C73787ad)] = 18;
//        decimalsMap[uint256(0xDc6dF65b2fA0322394a8af628Ad25Be7D7F413c2)] = 18;

    }

    /**
      * @notice set the precision
      * @dev function to update precision for an asset
      * @param newDecimals replacement oldDecimal
      */
    function setDecimals(uint256 newDecimals) public validCall{
        decimals = newDecimals;
    }


    /**
      * @notice Set prices in bulk
      * @dev function to update prices for an asset
      * @param prices replacement oldPrices
      */
    function setPrices(uint256[]memory assets,uint256[]memory prices) external validCall {
        require(assets.length == prices.length, "input arrays' length are not equal");
        uint256 len = assets.length;
        for (uint i=0;i<len;i++){
            priceMap[i] = prices[i];
        }
    }

    /**
      * @notice retrieve prices of assets in bulk
      * @dev function to get price for an assets
      * @param  assets Asset for which to get the price
      * @return uint mantissa of asset price (scaled by 1e8) or zero if unset or contract paused
      */
    function getPrices(uint256[]memory assets) public view returns (uint256[]memory) {
        uint256 len = assets.length;
        uint256[] memory prices = new uint256[](len);
        for (uint i=0;i<len;i++){
            prices[i] = getUnderlyingPrice(assets[i]);
        }
        return prices;
    }

    /**
      * @notice retrieves price of an asset
      * @dev function to get price for an asset
      * @param asset Asset for which to get the price
      * @return uint mantissa of asset price (scaled by 1e8) or zero if unset or contract paused
      */
    function getPrice(address asset) public view returns (uint256) {
        return getUnderlyingPrice(uint256(asset));
    }

    /**
      * @notice get price based on index
      * @dev function to get price for index
      * @param underlying for which to get the price
      * @return uint mantissa of asset price (scaled by 1e8) or zero if unset or contract paused
      */
    function getUnderlyingPrice(uint256 underlying) public view returns (uint256) {
        AggregatorV3Interface assetsPrice = assetsMap[underlying];
        if (address(assetsPrice) != address(0)){
            (, int price,,,) = assetsPrice.latestRoundData();
            uint256 tokenDecimals = decimalsMap[underlying];
            if (tokenDecimals < 18){
                return uint256(price)/decimals*(10**(18-tokenDecimals));
            }else if (tokenDecimals > 18){
                return uint256(price)/decimals/(10**(18-tokenDecimals));
            }else{
                return uint256(price)/decimals;
            }
        }else {
            return priceMap[underlying];
        }
    }


    /**
      * @notice set price of an asset
      * @dev function to set price for an asset
      * @param asset Asset for which to set the price
      * @param price the Asset's price
      */
    function setPrice(address asset,uint256 price) public validCall {
        priceMap[uint256(asset)] = price;
    }

    /**
      * @notice set price of an underlying
      * @dev function to set price for an underlying
      * @param underlying underlying for which to set the price
      * @param price the underlying's price
      */
    function setUnderlyingPrice(uint256 underlying,uint256 price) public validCall {
        require(underlying>0 , "underlying cannot be zero");
        priceMap[underlying] = price;
    }

    /**
      * @notice set price of an asset
      * @dev function to set price for an asset
      * @param asset Asset for which to set the price
      * @param aggergator the Asset's aggergator
      */
    function setAssetsAggregator(address asset,address aggergator,uint256 _decimals) public validCall {
        assetsMap[uint256(asset)] = AggregatorV3Interface(aggergator);
        decimalsMap[uint256(asset)] = _decimals;
    }

    /**
      * @notice set price of an underlying
      * @dev function to set price for an underlying
      * @param underlying underlying for which to set the price
      * @param aggergator the underlying's aggergator
      */
    function setUnderlyingAggregator(uint256 underlying,address aggergator,uint256 _decimals) public validCall {
        require(underlying>0 , "underlying cannot be zero");
        assetsMap[underlying] = AggregatorV3Interface(aggergator);
        decimalsMap[underlying] = _decimals;
    }

    /** @notice get asset aggregator based on asset
      * @dev function to get aggregator for asset
      * @param asset for which to get the aggregator
      * @ return  an asset aggregator
      */
    function getAssetsAggregator(address asset) public view returns (address,uint256) {
        return (address(assetsMap[uint256(asset)]),decimalsMap[uint256(asset)]);
    }

     /**
       * @notice get asset aggregator based on index
       * @dev function to get aggregator for index
       * @param underlying for which to get the aggregator
       * @ return an asset aggregator
       */
    function getUnderlyingAggregator(uint256 underlying) public view returns (address,uint256) {
        return (address(assetsMap[underlying]),decimalsMap[underlying]);
    }

}
