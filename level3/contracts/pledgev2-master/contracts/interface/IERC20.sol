// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

/**
 * @dev ERC20 代币标准接口
 * 参考: https://eips.ethereum.org/EIPS/eip-20
 */
interface IERC20 {
    /**
     * @dev 当代币被转移时触发的事件，包括零值转移
     * @param from 代币转出地址
     * @param to 代币转入地址
     * @param value 转移的代币数量
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev 当代币的授权额度被设置时触发的事件，包括零值授权
     * @param owner 代币持有者地址
     * @param spender 被授权的支出者地址
     * @param value 授权的代币数量
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);

    /**
     * @dev 返回代币的总供应量
     * @return 总供应量
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev 返回指定账户拥有的代币数量
     * @param account 要查询的账户地址
     * @return 账户拥有的代币数量
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev 转移指定数量的代币到目标地址
     * @param to 接收代币的目标地址
     * @param amount 要转移的代币数量
     * @return 如果转移成功则返回 true
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     * @dev 返回 spender 被允许从 owner 账户中支出的代币数量
     * @param owner 代币持有者地址
     * @param spender 被授权的支出者地址
     * @return 授权的代币数量
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev 设置对 spender 的授权额度为 amount
     * @param spender 被授权的支出者地址
     * @param amount 授权的代币数量
     * @return 如果授权成功则返回 true
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev 从 from 地址转移 amount 数量的代币到 to 地址，要求调用者有足够的授权额度
     * @param from 代币转出地址
     * @param to 代币转入地址
     * @param amount 要转移的代币数量
     * @return 如果转移成功则返回 true
     */
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
}// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP. Does not include
 * the optional functions; to access them see {ERC20Detailed}.
 */
interface IERC20 {
    function decimals() external view returns (uint8);
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `recipient`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address recipient, uint256 amount) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `sender` to `recipient` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);

      /**
     * EXTERNAL FUNCTION
     *
     * @dev change token name
     * @param _name token name
     * @param _symbol token symbol
     *
     */
    function changeTokenName(string calldata _name, string calldata _symbol)external;

    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);
}