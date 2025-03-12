const { ethers } = require("hardhat");

async function main() {
    const ONE_HOUR_IN_SECS = 60 * 60;
    const unlockTime = Math.floor(Date.now() / 1000) + ONE_HOUR_IN_SECS;

    const Lock = await ethers.getContractFactory("Lock");

    // 部署合约并发送 1 ETH
    const lock = await Lock.deploy(unlockTime, {
        value: ethers.parseEther("0.0001"),
    });

    // 等待部署确认
    await lock.waitForDeployment();

    // 获取合约地址（通过 .target 属性）
    console.log("Lock 合约地址:", lock.target);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });