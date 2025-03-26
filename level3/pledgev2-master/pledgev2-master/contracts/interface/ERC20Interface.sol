// SPDX-License-Identifier: MIT

pragma solidity 0.6.12;


interface ERC20Interface {
  //查询账户余额
  function balanceOf(address user) external view returns (uint256);
}
