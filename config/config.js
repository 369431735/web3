const fs = require('fs');
const path = require('path');
const yaml = require('yaml');

// 读取 YAML 配置文件
const configFile = fs.readFileSync(path.join(__dirname, '../../task2/config.yml'), 'utf8');
const config = yaml.parse(configFile);

module.exports = config; 