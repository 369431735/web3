async function main() {
    const initialSupply = 1000000; // 初始供应量 100万
    const MyToken = await ethers.getContractFactory("MyToken");
    const myToken = await MyToken.deploy(initialSupply);
    await myToken.waitForDeployment();
    console.log("合约地址:", await myToken.getAddress());
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});