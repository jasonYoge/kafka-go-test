package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var httpServer *gin.Engine

func init() {
	env := flag.String("e", "release", "Set running mode")
	flag.Parse()

	gin.SetMode(*env)
	httpServer = gin.New()
}

func postDataToKafka(ctx *gin.Context) {
	ctx.String(200, "running...")
}

func main() {
	httpServer.POST("/api/v1/data", postDataToKafka)

	httpServer.Run(":8001")
}