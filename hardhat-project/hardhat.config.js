require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

module.exports = {
  networks: {
    goerli: {
      url: `https://goerli.infura.io/v3/${process.env.INFURA_PROJECT_ID}`,
      accounts: [process.env.PRIVATE_KEY]
    }
  },
  solidity: "0.8.28",
  paths: {
    sources: "./contracts", // 确保路径指向合约文件夹
  }
};