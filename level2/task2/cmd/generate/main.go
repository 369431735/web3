package main

import (
	"log"

	"github.com/swaggo/swag"
	"github.com/swaggo/swag/cmd/swag"
)

func main() {
	// 设置 swag 命令
	swagCmd := swag.Command{
		Short: "生成 Swagger 文档",
		Long:  "生成 Swagger 文档，包括 docs.go、swagger.json 和 swagger.yaml",
	}

	// 执行 swag init 命令
	if err := swagCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
