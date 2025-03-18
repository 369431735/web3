const express = require('express');
const swaggerUi = require('swagger-ui-express');
const swaggerSpecs = require('./swagger');
const exampleRoutes = require('./routes/example');
const config = require('./config/config');

const app = express();
const { port, host } = config.server;

// 中间件
app.use(express.json());

// Swagger UI
app.use(config.swagger.path, swaggerUi.serve, swaggerUi.setup(swaggerSpecs));

// 路由
app.use('/api', exampleRoutes);

// 启动服务器
app.listen(port, host, () => {
  console.log(`服务器运行在 http://${host}:${port}`);
  console.log(`Swagger 文档可在 http://${host}:${port}${config.swagger.path} 访问`);
}); 