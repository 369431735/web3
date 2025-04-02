package controllers

import (
	"net/http"
	"pledge-backend/api/models/ws"
	"pledge-backend/log"
	"pledge-backend/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// PriceController 价格控制器
// 处理代币价格相关的API请求，主要通过WebSocket提供实时价格数据
type PriceController struct {
}

// NewPrice 提供PLGR代币实时价格的WebSocket服务
// 将HTTP连接升级为WebSocket连接，为客户端提供实时价格推送
// ctx: Gin上下文
func (c *PriceController) NewPrice(ctx *gin.Context) {

	// 添加异常恢复机制，防止panic导致服务中断
	defer func() {
		recoverRes := recover()
		if recoverRes != nil {
			log.Logger.Sugar().Error("new price recover ", recoverRes)
		}
	}()

	// 将HTTP连接升级为WebSocket连接
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:   1024,            // 读缓冲区大小
		WriteBufferSize:  1024,            // 写缓冲区大小
		HandshakeTimeout: 5 * time.Second, // 握手超时时间
		CheckOrigin: func(r *http.Request) bool { // 跨域检查
			return true // 允许所有来源的连接
		},
	}).Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Logger.Sugar().Error("websocket request err:", err)
		return
	}

	// 生成一个唯一的会话ID
	randomId := ""
	remoteIP, ok := ctx.RemoteIP()
	if ok {
		// 使用客户端IP地址和随机字符串生成会话ID
		randomId = strings.Replace(remoteIP.String(), ".", "_", -1) + "_" + utils.GetRandomString(23)
	} else {
		// 如果无法获取IP，则使用完全随机的字符串作为ID
		randomId = utils.GetRandomString(32)
	}

	// 创建WebSocket服务实例
	server := &ws.Server{
		Id:       randomId,               // 客户端唯一标识
		Socket:   conn,                   // WebSocket连接
		Send:     make(chan []byte, 800), // 发送消息的通道
		LastTime: time.Now().Unix(),      // 最后活跃时间
	}

	// 启动goroutine处理WebSocket的读写操作
	go server.ReadAndWrite()
}
