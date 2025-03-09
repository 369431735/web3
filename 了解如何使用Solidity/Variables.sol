// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
contract Variables{
    //申明在方法外的变量，储存在链上
    uint256  public count;


    function doSomething() public  {
        //申明在方法内部,方法结束后销毁
        uint256 i=20;

        //msg block全局变量
        uint256 timestamp = block.timestamp;
        address sender = msg.sender;
    }


}
