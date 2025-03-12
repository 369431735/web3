
const hre = require("hardhat");

async function main() {
    // 获取合约工厂
    const SimpleStorage = await hre.ethers.getContractFactory("./contracts/Lock");

    // 部署合约
    const simpleStorage = await SimpleStorage.deploy();

    // 等待合约部署完成
    await simpleStorage.deployed();

    console.log("SimpleStorage deployed to:", simpleStorage.address);
}

// 调用主函数并处理错误
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });