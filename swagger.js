const swaggerJsdoc = require('swagger-jsdoc');
const config = require('./config/config');

const options = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: config.swagger.title,
      version: config.swagger.version,
      description: config.swagger.description,
    },
    servers: [
      {
        url: `http://${config.server.host}:${config.server.port}`,
        description: 'Development server',
      },
    ],
  },
  apis: ['./routes/*.js'], // 指定包含 API 路由的文件路径
};

const specs = swaggerJsdoc(options);

module.exports = specs; 