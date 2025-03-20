const { ethers } = require("hardhat");
const mnemonic = "pioneer tiger vintage negative shove brisk hybrid tobacco shove dragon find volcano";

async function main() {
    // 使用 HDNodeWallet 生成钱包
    const wallet = ethers.HDNodeWallet.fromMnemonic(
        ethers.Mnemonic.fromPhrase(mnemonic),
        "m/44'/60'/0'/0/0" // 派生路径
    );

    console.log("Account #0 address:", wallet.address);
    console.log("Private key:", wallet.privateKey);
}

main().catch(console.error);