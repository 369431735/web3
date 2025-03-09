// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract IfElse  {
   function get(uint8  i) public returns  (uint8){
       if(i<10){
           return 1;
       }else if(i<20){
           return 2;
       }else{
           return 3;
       }
   }
    function gett(uint8  i) public returns  (uint8){
      return i>10?1:2;

    }

}
