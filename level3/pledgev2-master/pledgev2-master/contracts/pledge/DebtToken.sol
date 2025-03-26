// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./AddressPrivileges.sol";

/**
 * @dev 债务代币合约,用于管理借贷池中的债务代币(包括JP和SP代币)
 */
contract DebtToken is ERC20, AddressPrivileges {
    /**
     * @notice 构造函数
     * @param name_ 代币名称
     * @param symbol_ 代币符号
     * @param multiSignature 多签地址
     */
    constructor(
        string memory name_,
        string memory symbol_,
        address multiSignature
    ) public ERC20(name_, symbol_) AddressPrivileges(multiSignature) {}

    /**
     * @notice 铸造代币
     * @dev 只有铸造者可以调用此函数
     * @param account 接收代币的地址
     * @param amount 铸造的数量
     */
    function mint(address account, uint256 amount) external onlyMinter {
        _mint(account, amount);
    }

    /**
     * @notice 销毁代币
     * @dev 从调用者地址销毁指定数量的代币
     * @param amount 销毁的数量
     */
    function burn(address account, uint256 amount) external onlyMinter {
        _burn(account, amount);
    }
}