// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

interface IDebtToken {
     /**
     * @dev Returns the amount of tokens owned by `account`.
     返回指定账户拥有的代币数量
     */
    function balanceOf(address account) external view returns (uint256);

     /**
     * @dev Returns the amount of tokens in existence.
     返回代币的总供应量
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Minting tokens for specific accounts.
     为指定账户铸造（生成）指定数量的代币。
     */
    function mint(address account, uint256 amount) external;

     /**
     * @dev Burning tokens for specific accounts.
     为指定账户销毁（减少）指定数量的代币
     */
    function burn(address account, uint256 amount) external;

