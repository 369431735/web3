// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

interface IBscPledgeOracle {
    /**
      * @notice retrieves price of an asset
      * @dev function to get price for an asset
      * @param asset Asset for which to get the price
      * @return uint mantissa of asset price (scaled by 1e8) or zero if unset or contract paused
      */
    // 获取指定资产的价格
    function getPrice(address asset) external view returns (uint256);

    /**
      * @notice retrieves underlying price of a cToken
      * @dev function to get underlying price for a cToken
      * @param cToken cToken for which to get the underlying price
      * @return uint mantissa of underlying price (scaled by 1e8) or zero if unset or contract paused
      */
    // 获取指定 cToken 的基础价格
    function getUnderlyingPrice(uint256 cToken) external view returns (uint256);

    /**
      * @notice retrieves prices of multiple assets
      * @dev function to get prices for multiple assets
      * @param assets Array of assets for which to get the prices
      * @return uint[] Array of mantissas of asset prices (scaled by 1e8) or zero if unset or contract paused
      */
    // 获取多个资产的价格
    function getPrices(uint256[] calldata assets) external view returns (uint256[] memory);
}