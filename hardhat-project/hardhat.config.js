//启动本地网络 npx hardhat node

require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
const mnemonic = "pioneer tiger vintage negative shove brisk hybrid tobacco shove dragon find volcano"; // 替换为你的助记词

  module.exports = {
    solidity: "0.8.28",
    networks: {
      local: {
        url: "http://127.0.0.1:8545", // 本地网络 RPC 地址
        chainId: 1337, // 本地网络 ID
        accounts: {
          mnemonic: mnemonic, // 使用助记词生成账户
          path: "m/44'/60'/0'/0",
          initialIndex: 0,
          count: 20, // 生成 20 个测试账户
        },
      },
    },
  };