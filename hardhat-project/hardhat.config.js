require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
const mnemonic = "pioneer tiger vintage negative shove brisk hybrid tobacco shove dragon find volcano";

module.exports = {
  solidity: "0.8.28",
  networks: {
    // 内置的本地开发网络（显式启用WebSocket）
    hardhat: {
      chainId: 1337,
      // 配置HTTP和WebSocket服务
      http: {
        enabled: true,
        port: 8545
      },
      websockets: {
        enabled: true,   // 启用WebSocket
        port: 8545       // 使用与HTTP相同的端口
      },
      accounts: {
        mnemonic: mnemonic,
        path: "m/44'/60'/0'/0",
        initialIndex: 0,
        count: 20,
      }
    },
    // 本地节点别名（指向同一服务）
    localhost: {
      url: "http://127.0.0.1:8545",
      chainId: 1337
    }
  }
};

