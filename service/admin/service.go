package admin

import (
	"git.ezbuy.me/ezbuy/evtalk/common/ginrpc"

	"github.com/gin-gonic/gin"
)

func Restful() {
	router := gin.Default()
	router.Use(ginrpc.Cors())
	base := ginrpc.New("etadmin")
	base.Register(router, new(Admin))
	router.Run(":8081")
}
