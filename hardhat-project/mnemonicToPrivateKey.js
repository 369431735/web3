const { ethers } = require("ethers");

async function main() {
    const mnemonic = "pioneer tiger vintage negative shove brisk hybrid tobacco shove dragon find volcano"; // 替换为你的助记词
    const wallet = ethers.Wallet.fromMnemonic(mnemonic);
    console.log("Private Key:", wallet.privateKey);
    console.log("Address:", wallet.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });