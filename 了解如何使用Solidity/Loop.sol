// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract Loop {
    function loop() public {
        for(unit i ;i<10;i++){
            if(i==3){
                continue;
            }
            if(i==5){
                break;
            }
        }
    unit j;
        while(j<10){
            j++;
        }

    }



}
