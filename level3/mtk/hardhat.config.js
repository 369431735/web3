require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

module.exports = {
  solidity: "0.8.28",
  networks: {
    sepolia: {
      url: process.env.SEPOLIA_RPC_URL, // 使用 Infura 或 Alchemy 的 RPC
      accounts: [process.env.PRIVATE_KEY], // 你的钱包私钥
    }
  }
};