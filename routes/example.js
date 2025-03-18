const express = require('express');
const router = express.Router();

/**
 * @swagger
 * tags:
 *   name: 示例接口
 *   description: 用于演示的 API 接口集合
 */

/**
 * @swagger
 * /api/hello:
 *   get:
 *     tags: [示例接口]
 *     summary: 返回问候消息
 *     description: 一个简单的示例 API 端点，用于测试 API 是否正常工作
 *     operationId: getHello
 *     parameters:
 *       - in: query
 *         name: name
 *         schema:
 *           type: string
 *         description: 可选的用户名参数，如果不提供则使用默认问候语
 *         required: false
 *     responses:
 *       200:
 *         description: 成功返回问候消息
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 message:
 *                   type: string
 *                   description: 问候消息内容
 *                   example: Hello from Web3 API!
 *                 timestamp:
 *                   type: string
 *                   format: date-time
 *                   description: 响应生成的时间戳
 *                   example: "2024-03-18T10:00:00Z"
 *       400:
 *         description: 请求参数错误
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: string
 *                   description: 错误信息
 *                   example: "Invalid name parameter"
 */
router.get('/hello', (req, res) => {
  const { name } = req.query;
  const message = name ? `Hello ${name} from Web3 API!` : 'Hello from Web3 API!';
  res.json({
    message,
    timestamp: new Date().toISOString()
  });
});

/**
 * @swagger
 * /api/status:
 *   get:
 *     tags: [示例接口]
 *     summary: 获取系统状态
 *     description: 返回当前系统的运行状态信息
 *     operationId: getStatus
 *     responses:
 *       200:
 *         description: 成功返回系统状态信息
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 status:
 *                   type: string
 *                   description: 系统运行状态
 *                   enum: [running, maintenance, error]
 *                   example: running
 *                 uptime:
 *                   type: number
 *                   description: 系统运行时间（秒）
 *                   example: 3600
 *                 version:
 *                   type: string
 *                   description: API 版本号
 *                   example: "1.0.0"
 *                 lastCheck:
 *                   type: string
 *                   format: date-time
 *                   description: 最后一次状态检查时间
 *                   example: "2024-03-18T10:00:00Z"
 */
router.get('/status', (req, res) => {
  res.json({
    status: 'running',
    uptime: process.uptime(),
    version: '1.0.0',
    lastCheck: new Date().toISOString()
  });
});

/**
 * @swagger
 * /api/echo:
 *   post:
 *     tags: [示例接口]
 *     summary: 回显请求数据
 *     description: 接收 POST 请求中的数据并返回，用于测试 POST 请求
 *     operationId: postEcho
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               message:
 *                 type: string
 *                 description: 要回显的消息
 *                 example: "Hello World"
 *               timestamp:
 *                 type: string
 *                 format: date-time
 *                 description: 可选的时间戳
 *                 example: "2024-03-18T10:00:00Z"
 *     responses:
 *       200:
 *         description: 成功返回回显数据
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 received:
 *                   type: object
 *                   description: 接收到的请求数据
 *                 serverTime:
 *                   type: string
 *                   format: date-time
 *                   description: 服务器处理时间
 *                   example: "2024-03-18T10:00:00Z"
 *       400:
 *         description: 请求体格式错误
 *         content:
 *           application/json:
 *             schema:
 *               type: object
 *               properties:
 *                 error:
 *                   type: string
 *                   description: 错误信息
 *                   example: "Invalid request body"
 */
router.post('/echo', (req, res) => {
  res.json({
    received: req.body,
    serverTime: new Date().toISOString()
  });
});

module.exports = router; 