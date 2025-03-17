const hre = require("hardhat");
require('dotenv').config();

async function main() {
  console.log("开始部署合约...");

  // 获取部署账号
  const [deployer] = await hre.ethers.getSigners();
  console.log("使用账号部署:", deployer.address);
  console.log("账号余额:", hre.ethers.formatEther(await deployer.provider.getBalance(deployer.address)), "ETH");

  // 部署 Lock 合约
  // 功能：时间锁定合约，允许在指定时间后提取存入的ETH
  // 参数：unlockTime - 解锁时间（unix时间戳）
  // 初始值：0.1 ETH
  const currentTimestampInSeconds = Math.round(Date.now() / 1000);
  const unlockTime = currentTimestampInSeconds + 60; // 1分钟后解锁
  const Lock = await hre.ethers.getContractFactory("Lock");
  const lock = await Lock.deploy(unlockTime, { value: hre.ethers.parseEther("0.1") });
  await lock.waitForDeployment();
  console.log(`Lock 合约已部署到: ${await lock.getAddress()}`);

  // 部署 ERC20MinerReward 合约
  // 功能：矿工奖励代币合约，实现ERC20标准
  // 特点：每次挖矿可获得20个代币奖励
  const ERC20MinerReward = await hre.ethers.getContractFactory("ERC20MinerReward");
  const minerReward = await ERC20MinerReward.deploy();
  await minerReward.waitForDeployment();
  console.log(`ERC20MinerReward 合约已部署到: ${await minerReward.getAddress()}`);

  // 部署 Shipping 合约
  // 功能：物流跟踪合约，记录货物运输状态
  // 状态：Pending（待发货）、Shipped（已发货）、Delivered（已送达）
  const Shipping = await hre.ethers.getContractFactory("Shipping");
  const shipping = await Shipping.deploy();
  await shipping.waitForDeployment();
  console.log(`Shipping 合约已部署到: ${await shipping.getAddress()}`);

  // 部署 SimpleStorage 合约
  // 功能：简单存储合约，演示基本的状态变量存储
  // 操作：可以存储和读取一个数值
  const SimpleStorage = await hre.ethers.getContractFactory("SimpleStorage");
  const simpleStorage = await SimpleStorage.deploy();
  await simpleStorage.waitForDeployment();
  console.log(`SimpleStorage 合约已部署到: ${await simpleStorage.getAddress()}`);

  // 部署 Lottery 合约
  // 功能：去中心化彩票系统
  // 特点：
  // - 最小参与金额：0.01 ETH
  // - 自动生成随机数选择获胜者
  // - 包含冷却时间机制
  const Lottery = await hre.ethers.getContractFactory("Lottery");
  const lottery = await Lottery.deploy();
  await lottery.waitForDeployment();
  console.log(`Lottery 合约已部署到: ${await lottery.getAddress()}`);

  // 部署 SimpleAuction 合约
  // 功能：简单的拍卖系统
  // 特点：
  // - 设定拍卖时间
  // - 自动处理最高出价
  // - 允许撤回非最高出价
  const biddingTime = 3600; // 1小时的拍卖时间
  const SimpleAuction = await hre.ethers.getContractFactory("SimpleAuction");
  const simpleAuction = await SimpleAuction.deploy(biddingTime, deployer.address);
  await simpleAuction.waitForDeployment();
  console.log(`SimpleAuction 合约已部署到: ${await simpleAuction.getAddress()}`);

  // 部署 Purchase 合约
  // 功能：安全购买合约
  // 特点：
  // - 买家卖家双方托管
  // - 自动处理退款
  // - 确认收货后自动结算
  // 初始值：1.0 ETH
  const Purchase = await hre.ethers.getContractFactory("Purchase");
  const purchase = await Purchase.deploy({ value: hre.ethers.parseEther("1.0") });
  await purchase.waitForDeployment();
  console.log(`Purchase 合约已部署到: ${await purchase.getAddress()}`);

  // 部署 ArrayDemo 合约
  // 功能：数组操作示例合约
  // 演示：动态数组的基本操作（添加、删除、修改、查询）
  const ArrayDemo = await hre.ethers.getContractFactory("arrayDemo");
  const arrayDemo = await ArrayDemo.deploy();
  await arrayDemo.waitForDeployment();
  console.log(`ArrayDemo 合约已部署到: ${await arrayDemo.getAddress()}`);

  // 部署 Ballot 合约
  // 功能：投票系统
  // 特点：
  // - 支持多个提案
  // - 每个投票人只能投票一次
  // - 可以委托投票权
  const proposalNames = ["Proposal1", "Proposal2", "Proposal3"].map(name => 
    ethers.encodeBytes32String(name)
  );
  const Ballot = await hre.ethers.getContractFactory("Ballot");
  const ballot = await Ballot.deploy(proposalNames);
  await ballot.waitForDeployment();
  console.log(`Ballot 合约已部署到: ${await ballot.getAddress()}`);

  // 部署 Index 合约
  // 功能：基础合约示例
  // 演示：基本的合约结构和状态变量使用
  const Index = await hre.ethers.getContractFactory("Index");
  const index = await Index.deploy();
  await index.waitForDeployment();
  console.log(`Index 合约已部署到: ${await index.getAddress()}`);

  // 打印所有合约地址的汇总
  console.log("\n合约部署汇总:");
  console.log("====================");
  console.log(`Lock: ${await lock.getAddress()}`);
  console.log(`ERC20MinerReward: ${await minerReward.getAddress()}`);
  console.log(`Shipping: ${await shipping.getAddress()}`);
  console.log(`SimpleStorage: ${await simpleStorage.getAddress()}`);
  console.log(`Lottery: ${await lottery.getAddress()}`);
  console.log(`SimpleAuction: ${await simpleAuction.getAddress()}`);
  console.log(`Purchase: ${await purchase.getAddress()}`);
  console.log(`ArrayDemo: ${await arrayDemo.getAddress()}`);
  console.log(`Ballot: ${await ballot.getAddress()}`);
  console.log(`Index: ${await index.getAddress()}`);
  console.log("====================\n");

  console.log("所有合约部署完成！");
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});