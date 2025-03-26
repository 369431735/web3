// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "../library/SafeMathExt.sol";
import "../multiSignature/multiSignatureClient.sol";

contract AddressPrivileges is multiSignatureClient {
    using SafeMathExt for uint256;

    // 地址类型的枚举
    enum AddressType{ NONE, OPERATOR, MINTER }
    // 默认选择为NONE类型
    AddressType constant defaultChoice = AddressType.NONE;

    // 存储地址特权信息映射：address => AddressType
    mapping(address => AddressType) private addressPrivileges;

    /**
     * @notice 构造函数，初始化多重签名地址
     * @param multiSignature 多重签名合约地址
     */
    constructor(address multiSignature) multiSignatureClient(multiSignature) public {
    }

    /**
     * @notice 设置操作员权限
     * @dev 只能通过多重签名调用
     * @param _address 要设置权限的地址
     * @param _enable 是否启用权限
     */
    function setOperator(address _address, bool _enable) external validCall {
        require(_address != address(0), "AddressPrivileges: zero address");
        if (_enable) {
            addressPrivileges[_address] = AddressType.OPERATOR;
        } else {
            addressPrivileges[_address] = AddressType.NONE;
        }
    }

    /**
     * @notice 设置铸币权限
     * @dev 只能通过多重签名调用
     * @param _address 要设置权限的地址
     * @param _enable 是否启用权限
     */
    function setMinter(address _address, bool _enable) external validCall {
        require(_address != address(0), "AddressPrivileges: zero address");
        if (_enable) {
            addressPrivileges[_address] = AddressType.MINTER;
        } else {
            addressPrivileges[_address] = AddressType.NONE;
        }
    }

    /**
     * @notice 检查地址是否为操作员
     * @param _address 要检查的地址
     * @return bool 是否为操作员
     */
    function isOperator(address _address) public view returns (bool) {
        return addressPrivileges[_address] == AddressType.OPERATOR;
    }

    /**
     * @notice 检查地址是否为铸币者
     * @param _address 要检查的地址
     * @return bool 是否为铸币者
     */
    function isMinter(address _address) public view returns (bool) {
        return addressPrivileges[_address] == AddressType.MINTER;
    }

    /**
     * @notice 要求调用者必须是操作员
     */
    modifier onlyOperator() {
        require(isOperator(msg.sender), "AddressPrivileges: caller is not operator");
        _;
    }

    /**
     * @notice 要求调用者必须是铸币者
     */
    modifier onlyMinter() {
        require(isMinter(msg.sender), "AddressPrivileges: caller is not minter");
        _;
    }
}// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;

import "../multiSignature/multiSignatureClient.sol";
import "@openzeppelin/contracts/utils/EnumerableSet.sol";

/**
 * @dev Collection of functions related to the address type
 */
contract AddressPrivileges is multiSignatureClient {

    constructor(address multiSignature) multiSignatureClient(multiSignature) public {
    }

    using EnumerableSet for EnumerableSet.AddressSet;
    EnumerableSet.AddressSet private _minters;

    /**
      * @notice add a minter
      * @dev function to add a minter for an asset
      * @param _addMinter add a  minter address
      # @return true or false
      */
    function addMinter(address _addMinter) public validCall returns (bool) {
        require(_addMinter != address(0), "Token: _addMinter is the zero address");
        return EnumerableSet.add(_minters, _addMinter);
    }

    /**
      * @notice delete a minter
      * @dev function to delete a minter for an asset
      * @param _delMinter delete a minter address
      # @return true or false
      */
    function delMinter(address _delMinter) public validCall returns (bool) {
        require(_delMinter != address(0), "Token: _delMinter is the zero address");
        return EnumerableSet.remove(_minters, _delMinter);
    }

    /**
      * @notice get minter list length
      * @dev function to get the minter list length
      # @return the lenght of minter list
      */
    function getMinterLength() public view returns (uint256) {
        return EnumerableSet.length(_minters);
    }

    /**
      * @notice Determine if this address is minter
      * @dev function to judge address
      * @param account is a condition
      # @return true or false
      */
    function isMinter(address account) public view returns (bool) {
        return EnumerableSet.contains(_minters, account);
    }

     /**
      * @notice Get minter account according to index
      * @dev function to get minter account
      * @param _index of index
      # @return  a minter account
      */
    function getMinter(uint256 _index) public view  returns (address){
        require(_index <= getMinterLength() - 1, "Token: index out of bounds");
        return EnumerableSet.at(_minters, _index);
    }

    modifier onlyMinter() {
        require(isMinter(msg.sender), "Token: caller is not the minter");
        _;
    }

}