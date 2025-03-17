const fs = require('fs');
const path = require('path');

async function main() {
    // 创建备份目录
    const backupDir = path.join(__dirname, '../contracts_backup');
    if (!fs.existsSync(backupDir)) {
        fs.mkdirSync(backupDir);
    }

    // 复制合约文件到备份目录
    const contractsDir = path.join(__dirname, '../contracts');
    const files = fs.readdirSync(contractsDir);
    
    console.log('开始备份合约文件...');
    for (const file of files) {
        if (file.endsWith('.sol')) {
            const sourcePath = path.join(contractsDir, file);
            const targetPath = path.join(backupDir, file);
            fs.copyFileSync(sourcePath, targetPath);
            console.log(`已备份: ${file}`);
        }
    }

    // 编译合约生成ABI
    console.log('\n开始编译合约生成ABI...');
    await hre.run('compile');
    
    console.log('\n完成！');
    console.log(`合约文件备份位置: ${backupDir}`);
    console.log('ABI文件位置: artifacts/contracts/');
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    }); 