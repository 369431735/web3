// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract SimpleStorage {


    uint256 public test;

   public change(uint256 i){
     test=i;
  }

  public get() returns (uint256){
   return test;
   }
}
