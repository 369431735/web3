// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

/**
 * @dev Collection of functions related to the address type
 */
library Address {
    /**
     * @dev Returns true if `account` is a contract.
     *
     * [IMPORTANT]
     * ====
     * It is unsafe to assume that an address for which this function returns
     * false is an externally-owned account (EOA) and not a contract.
     *
     * Among others, `isContract` will return false for the following
     * types of addresses:
     *
     *  - an externally-owned account
     *  - a contract in construction
     *  - an address where a contract will be created
     *  - an address where a contract lived, but was destroyed
     * ====
     */
    function isContract(address account) internal view returns (bool) {
        // This method relies on extcodesize, which returns 0 for contracts in
        // construction, since the code is only stored at the end of the
        // constructor execution.

        uint256 size;
        // solhint-disable-next-line no-inline-assembly
        assembly { size := extcodesize(account) }
        return size > 0;
    }

    /**
     * @dev Replacement for Solidity's `transfer`: sends `amount` wei to
     * `recipient`, forwarding all available gas and reverting on errors.
     *
     * https://eips.ethereum.org/EIPS/eip-1884[EIP1884] increases the gas cost
     * of certain opcodes, possibly making contracts go over the 2300 gas limit
     * imposed by `transfer`, making them unable to receive funds via
     * `transfer`. {sendValue} removes this limitation.
     *
     * https://diligence.consensys.net/posts/2019/09/stop-using-soliditys-transfer-now/[Learn more].
     *
     * IMPORTANT: because control is transferred to `recipient`, care must be
     * taken to not create reentrancy vulnerabilities. Consider using
     * {ReentrancyGuard} or the
     * https://solidity.readthedocs.io/en/v0.5.11/security-considerations.html#use-the-checks-effects-interactions-pattern[checks-effects-interactions pattern].
     */
    function sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        // solhint-disable-next-line avoid-low-level-calls, avoid-call-value
        (bool success, ) = recipient.call{value:amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }

    /**
     * @dev Performs a Solidity function call using a low level `call`. A
     * plain`call` is an unsafe replacement for a function call: use this
     * function instead.
     *
     * If `target` reverts with a revert reason, it is bubbled up by this
     * function (like regular Solidity function calls).
     *
     * Returns the raw returned data. To convert to the expected return value,
     * use https://solidity.readthedocs.io/en/latest/units-and-global-variables.html?highlight=abi.decode#abi-encoding-and-decoding-functions[`abi.decode`].
     *
     * Requirements:
     *
     * - `target` must be a contract.
     * - calling `target` with `data` must not revert.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data) internal returns (bytes memory) {
      return functionCall(target, data, "Address: low-level call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`], but with
     * `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCall(address target, bytes memory data, string memory errorMessage) internal returns (bytes memory) {
        return functionCallWithValue(target, data, 0, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but also transferring `value` wei to `target`.
     *
     * Requirements:
     *
     * - the calling contract must have an ETH balance of at least `value`.
     * - the called Solidity function must be `payable`.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(address target, bytes memory data, uint256 value) internal returns (bytes memory) {
        return functionCallWithValue(target, data, value, "Address: low-level call with value failed");
    }

    /**
     * @dev Same as {xref-Address-functionCallWithValue-address-bytes-uint256-}[`functionCallWithValue`], but
     * with `errorMessage` as a fallback revert reason when `target` reverts.
     *
     * _Available since v3.1._
     */
    function functionCallWithValue(address target, bytes memory data, uint256 value, string memory errorMessage) internal returns (bytes memory) {
        require(address(this).balance >= value, "Address: insufficient balance for call");
        require(isContract(target), "Address: call to non-contract");

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory returndata) = target.call{value:value}(data);
        return _verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
        return functionStaticCall(target, data, "Address: low-level static call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a static call.
     *
     * _Available since v3.3._
     */
    function functionStaticCall(address target, bytes memory data, string memory errorMessage) internal view returns (bytes memory) {
        require(isContract(target), "Address: static call to non-contract");

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory returndata) = target.staticcall(data);
        return _verifyCallResult(success, returndata, errorMessage);
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-}[`functionCall`],
     * but performing a delegate call.
     *
     * _Available since v3.3._
     */
    function functionDelegateCall(address target, bytes memory data) internal returns (bytes memory) {
        return functionDelegateCall(target, data, "Address: low-level delegate call failed");
    }

    /**
     * @dev Same as {xref-Address-functionCall-address-bytes-string-}[`functionCall`],
     * but performing a delegate call.
     *// SPDX-License-Identifier: MIT

     pragma solidity 0.6.12;

     /**
      * @dev 与地址类型相关的函数集合
      */
     library Address {
         /**
          * @dev 检查一个地址是否是合约地址
          *
          * [重要提示]
          * ====
          * 不能假设此函数返回 false 的地址就一定是外部账户(EOA)而不是合约。
          *
          * 以下情况该函数会返回 false：
          *  - 外部拥有的账户地址
          *  - 正在构建中的合约
          *  - 将要创建合约的地址
          *  - 曾经存在但已被销毁的合约地址
          * ====
          */
         function isContract(address account) internal view returns (bool) {
             // 此方法依赖于 extcodesize，对于正在构建的合约将返回 0
             // 因为代码只在构造函数执行结束时才会存储
             uint256 size;
             // solhint-disable-next-line no-inline-assembly
             assembly { size := extcodesize(account) }
             return size > 0;
         }

         /**
          * @dev Solidity 的 transfer 函数的替代品：发送 amount 数量的 wei 到
          * recipient，转发所有可用的 gas 并在发生错误时回滚
          *
          * 这个函数移除了 EIP1884 引入的 2300 gas 限制
          *
          * 重要：由于控制权转移到了接收者，必须注意防止重入攻击。
          * 建议使用 {ReentrancyGuard} 或检查-生效-交互模式
          */
         function sendValue(address payable recipient, uint256 amount) internal {
             require(address(this).balance >= amount, "Address: 余额不足");

             (bool success, ) = recipient.call{value:amount}("");
             require(success, "Address: 发送失败，接收者可能已回滚");
         }

         /**
          * @dev 执行带有低级 call 的 Solidity 函数调用
          * 如果目标合约回滚并提供原因，该原因会被向上传递
          * 返回原始返回数据，使用 abi.decode 转换为期望的返回值
          *
          * 要求：
          * - target 必须是合约
          * - 调用 target 时不能回滚
          */
         function functionCall(address target, bytes memory data) internal returns (bytes memory) {
             return functionCall(target, data, "Address: 低级调用失败");
         }

         /**
          * @dev 与 functionCall 相同，但可以自定义错误消息
          */
         function functionCall(address target, bytes memory data, string memory errorMessage) internal returns (bytes memory) {
             return functionCallWithValue(target, data, 0, errorMessage);
         }

         /**
          * @dev 与 functionCall 相同，但同时转账 value 数量的 wei 到目标地址
          * 要求：
          * - 调用合约必须有足够的 ETH 余额
          * - 被调用的函数必须是 payable
          */
         function functionCallWithValue(address target, bytes memory data, uint256 value) internal returns (bytes memory) {
             return functionCallWithValue(target, data, value, "Address: 带值的低级调用失败");
         }

         /**
          * @dev 与 functionCallWithValue 相同，但可以自定义错误消息
          */
         function functionCallWithValue(address target, bytes memory data, uint256 value, string memory errorMessage) internal returns (bytes memory) {
             require(address(this).balance >= value, "Address: 调用所需余额不足");
             require(isContract(target), "Address: 调用非合约地址");

             (bool success, bytes memory returndata) = target.call{value:value}(data);
             return _verifyCallResult(success, returndata, errorMessage);
         }

         /**
          * @dev 与 functionCall 相同，但执行 staticcall（不修改状态的调用）
          */
         function functionStaticCall(address target, bytes memory data) internal view returns (bytes memory) {
             return functionStaticCall(target, data, "Address: 低级静态调用失败");
         }

         function functionStaticCall(address target, bytes memory data, string memory errorMessage) internal view returns (bytes memory) {
             require(isContract(target), "Address: 静态调用非合约地址");

             (bool success, bytes memory returndata) = target.staticcall(data);
             return _verifyCallResult(success, returndata, errorMessage);
         }

         /**
          * @dev 与 functionCall 相同，但执行 delegatecall（使用调用者的存储的调用）
          */
         function functionDelegateCall(address target, bytes memory data) internal returns (bytes memory) {
             return functionDelegateCall(target, data, "Address: 低级委托调用失败");
         }

         function functionDelegateCall(address target, bytes memory data, string memory errorMessage) internal returns (bytes memory) {
             require(isContract(target), "Address: 委托调用非合约地址");

             (bool success, bytes memory returndata) = target.delegatecall(data);
             return _verifyCallResult(success, returndata, errorMessage);
         }

         function _verifyCallResult(bool success, bytes memory returndata, string memory errorMessage) private pure returns(bytes memory) {
             if (success) {
                 return returndata;
             } else {
                 // 查找回滚原因并向上传递
                 if (returndata.length > 0) {
                     assembly {
                         let returndata_size := mload(returndata)
                         revert(add(32, returndata), returndata_size)
                     }
                 } else {
                     revert(errorMessage);
                 }
             }
         }
     }
     * _Available since v3.3._
     */
    function functionDelegateCall(address target, bytes memory data, string memory errorMessage) internal returns (bytes memory) {
        require(isContract(target), "Address: delegate call to non-contract");

        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory returndata) = target.delegatecall(data);
        return _verifyCallResult(success, returndata, errorMessage);
    }

    function _verifyCallResult(bool success, bytes memory returndata, string memory errorMessage) private pure returns(bytes memory) {
        if (success) {
            return returndata;
        } else {
            // Look for revert reason and bubble it up if present
            if (returndata.length > 0) {
                // The easiest way to bubble the revert reason is using memory via assembly

                // solhint-disable-next-line no-inline-assembly
                assembly {
                    let returndata_size := mload(returndata)
                    revert(add(32, returndata), returndata_size)
                }
            } else {
                revert(errorMessage);
            }
        }
    }
}