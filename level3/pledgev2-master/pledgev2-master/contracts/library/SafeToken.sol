// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "./SafeMath.sol";
import "./Address.sol";
import "../interface/IERC20.sol";

/**
 * @title SafeToken
 * @dev 用于安全地与 ERC20 代币进行交互的库
 */
library SafeToken {
    using SafeMath for uint256;
    using Address for address;

    /**
     * @notice 安全地将代币从发送者转移到接收者
     * @dev 调用代币合约的 transfer 函数
     * @param token ERC20 代币合约地址
     * @param to 接收者地址
     * @param value 要转移的代币数量
     * @return success 转账是否成功
     */
    function safeTransfer(address token, address to, uint256 value) internal returns (bool success) {
        (success,) = token.call(abi.encodeWithSelector(IERC20.transfer.selector, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), "SafeToken: transfer 失败");
    }

    /**
     * @notice 使用授权额度安全地将代币从一个地址转移到另一个地址
     * @dev 调用代币合约的 transferFrom 函数
     * @param token ERC20 代币合约地址
     * @param from 转出地址
     * @param to 接收地址
     * @param value 要转移的代币数量
     * @return success 转账是否成功
     */
    function safeTransferFrom(address token, address from, address to, uint256 value) internal returns (bool success) {
        (success,) = token.call(abi.encodeWithSelector(IERC20.transferFrom.selector, from, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), "SafeToken: transferFrom 失败");
    }

    /**
     * @notice 安全地授权支出者花费代币
     * @dev 调用代币合约的 approve 函数
     * @param token ERC20 代币合约地址
     * @param spender 被授权的支出者地址
     * @param value 授权的代币数量
     * @return success 授权是否成功
     */
    function safeApprove(address token, address spender, uint256 value) internal returns (bool success) {
        (success,) = token.call(abi.encodeWithSelector(IERC20.approve.selector, spender, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), "SafeToken: approve 失败");
    }

    /**
     * @notice 安全地获取账户的代币余额
     * @dev 调用代币合约的 balanceOf 函数
     * @param token ERC20 代币合约地址
     * @param account 要查询余额的账户地址
     * @return balance 账户的代币余额
     */
    function safeBalanceOf(address token, address account) internal view returns (uint256 balance) {
        (bool success, bytes memory data) = token.staticcall(abi.encodeWithSelector(IERC20.balanceOf.selector, account));
        require(success && data.length >= 32, "SafeToken: balanceOf 失败");
        balance = abi.decode(data, (uint256));
    }

    /**
     * @notice 安全地获取支出者的授权额度
     * @dev 调用代币合约的 allowance 函数
     * @param token ERC20 代币合约地址
     * @param owner 代币所有者地址
     * @param spender 被授权的支出者地址
     * @return allowance 授权额度
     */
    function safeAllowance(address token, address owner, address spender) internal view returns (uint256 allowance) {
        (bool success, bytes memory data) = token.staticcall(abi.encodeWithSelector(IERC20.allowance.selector, owner, spender));
        require(success && data.length >= 32, "SafeToken: allowance 失败");
        allowance = abi.decode(data, (uint256));
    }
}/**
 * @notice 安全地转移代币
 * @dev 调用代币合约的 transfer 函数
 * @param token ERC20 代币合约地址
 * @param to 接收代币的地址
 * @param value 转账金额
 * @return success 转账是否成功
 */
function safeTransfer(address token, address to, uint256 value) internal returns (bool success) {
    (bool success, bytes memory data) = token.call(abi.encodeWithSelector(IERC20.transfer.selector, to, value));
    require(success && (data.length == 0 || abi.decode(data, (bool))), "SafeToken: transfer 失败");
}

/**
 * @notice 安全地从指定账户转移代币
 * @dev 调用代币合约的 transferFrom 函数
 * @param token ERC20 代币合约地址
 * @param from 转出代币的地址
 * @param to 接收代币的地址
 * @param value 转账金额
 * @return success 转账是否成功
 */
function safeTransferFrom(address token, address from, address to, uint256 value) internal returns (bool success) {
    (bool success, bytes memory data) = token.call(abi.encodeWithSelector(IERC20.transferFrom.selector, from, to, value));
    require(success && (data.length == 0 || abi.decode(data, (bool))), "SafeToken: transferFrom 失败");
}// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;


import "../interface/ERC20Interface.sol";

library SafeToken {
  function myBalance(address token) internal view returns (uint256) {
    return ERC20Interface(token).balanceOf(address(this));
  }

  function balanceOf(address token, address user) internal view returns (uint256) {
    return ERC20Interface(token).balanceOf(user);
  }

  function safeApprove(address token, address to, uint256 value) internal {
    // bytes4(keccak256(bytes('approve(address,uint256)')));
    (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x095ea7b3, to, value));
    require(success && (data.length == 0 || abi.decode(data, (bool))), "!safeApprove");
  }

  function safeTransfer(address token, address to, uint256 value) internal {
    // bytes4(keccak256(bytes('transfer(address,uint256)')));
    (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0xa9059cbb, to, value));
    require(success && (data.length == 0 || abi.decode(data, (bool))), "!safeTransfer");
  }

  function safeTransferFrom(address token, address from, address to, uint256 value) internal {
    // bytes4(keccak256(bytes('transferFrom(address,address,uint256)')));
    (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x23b872dd, from, to, value));
    require(success && (data.length == 0 || abi.decode(data, (bool))), "!safeTransferFrom");
  }


}