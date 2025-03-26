// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "./SafeToken.sol";
import "./SafeMath.sol";
import "./Address.sol";
import "../interface/IERC20.sol";

/**
 * @title SafeTransfer
 * @dev 提供安全的代币转账和批准操作的库
 */
library SafeTransfer {
    using SafeMath for uint256;
    using Address for address;

    /**
     * @dev 安全转账指定数量的代币
     * @param token 代币合约地址
     * @param to 接收者地址
     * @param value 转账数量
     * @return success 转账是否成功
     */
    function safeTransfer(
        address token,
        address to,
        uint256 value
    ) internal returns (bool success) {
        (success,) = token.call(abi.encodeWithSelector(IERC20.transfer.selector, to, value));
        require(success, "SafeTransfer: transfer failed");
        return success;
    }

    /**
     * @dev 安全地从指定地址转账代币到目标地址
     * @param token 代币合约地址
     * @param from 发送者地址
     * @param to 接收者地址
     * @param value 转账数量
     * @return success 转账是否成功
     */
    function safeTransferFrom(
        address token,
        address from,
        address to,
        uint256 value
    ) internal returns (bool success) {
        (success,) = token.call(abi.encodeWithSelector(IERC20.transferFrom.selector, from, to, value));
        require(success, "SafeTransfer: transferFrom failed");
        return success;
    }

    /**
     * @dev 安全地批准指定地址使用代币
     * @param token 代币合约地址
     * @param spender 被批准的地址
     * @param value 批准的数量
     * @return success 批准是否成功
     */
    function safeApprove(
        address token,
        address spender,
        uint256 value
    ) internal returns (bool success) {
        (success,) = token.call(abi.encodeWithSelector(IERC20.approve.selector, spender, value));
        require(success, "SafeTransfer: approve failed");
        return success;
    }

    /**
     * @dev 获取代币余额
     * @param token 代币合约地址
     * @param account 查询的账户地址
     * @return balance 账户的代币余额
     */
    function safeBalanceOf(
        address token,
        address account
    ) internal view returns (uint256 balance) {
        (bool success, bytes memory data) = token.staticcall(abi.encodeWithSelector(IERC20.balanceOf.selector, account));
        require(success && data.length >= 32, "SafeTransfer: balanceOf failed");
        balance = abi.decode(data, (uint256));
    }

    /**
     * @dev 获取代币的授权额度
     * @param token 代币合约地址
     * @param owner 授权者地址
     * @param spender 被授权者地址
     * @return allowance 授权额度
     */
    function safeAllowance(
        address token,
        address owner,
        address spender
    ) internal view returns (uint256 allowance) {
        (bool success, bytes memory data) = token.staticcall(abi.encodeWithSelector(IERC20.allowance.selector, owner, spender));
        require(success && data.length >= 32, "SafeTransfer: allowance failed");
        allowance = abi.decode(data, (uint256));
    }
}
// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "./SafeErc20.sol";

contract SafeTransfer{

    using SafeERC20 for IERC20;
    event Redeem(address indexed recieptor,address indexed token,uint256 amount);

    /**
     * @notice  transfers money to the pool
     * @dev function to transfer
     * @param token of address
     * @param amount of amount
     * @return return amount
     */
    function getPayableAmount(address token,uint256 amount) internal returns (uint256) {
        if (token == address(0)){
            amount = msg.value;
        }else if (amount > 0){
            IERC20 oToken = IERC20(token);
            oToken.safeTransferFrom(msg.sender, address(this), amount);
        }
        return amount;
    }

    /**
     * @dev An auxiliary foundation which transter amount stake coins to recieptor.
     * @param recieptor account.
     * @param token address
     * @param amount redeem amount.
     */
    function _redeem(address payable recieptor,address token,uint256 amount) internal{
        if (token == address(0)){
            recieptor.transfer(amount);
        }else{
            IERC20 oToken = IERC20(token);
            oToken.safeTransfer(recieptor,amount);
        }
        emit Redeem(recieptor,token,amount);
    }
}