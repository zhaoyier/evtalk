package webapi

import (
	"git.ezbuy.me/ezbuy/evtalk/common/ginrpc"

	"github.com/gin-gonic/gin"
)

func Restful() {
	router := gin.Default()
	router.Use(ginrpc.Cors())
	base := ginrpc.New("evtalk")
	base.Register(router, new(Hello))
	router.Run(":8080")
}
