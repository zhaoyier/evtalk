package admin

import (
	rpc "git.ezbuy.me/ezbuy/evtalk/rpc/proto/etadmin"
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

//登录
func (a *Admin) AdminLogin(c *gin.Context, req *rpc.AdminLoginReq) (*rpc.AdminLoginResp, error) {
	c.SetCookie("evtalk", "token_002", 3600, "/", c.Request.Host, false, true)
	return &rpc.AdminLoginResp{
		Status:  true,
		Code:    "00000000",
		Message: "success",
		Data: &rpc.LoginData{
			TimeToExpire: 604800,
		},
	}, nil
}

//退出
func (a *Admin) AdminLogout(c *gin.Context, req *rpc.AdminLogoutReq) (*rpc.AdminLogoutResp, error) {
	return nil, nil
}
