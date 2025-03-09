// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract SimpleStorage {


    uint256 public test;

    function change(uint256 i) public{
     test=i;
  }

    function get() public returns (uint256){
   return test;
   }
}
