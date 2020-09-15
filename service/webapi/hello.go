package webapi

import (
	// "git.ezbuy.me/ezbuy/evtalk/rpc/webapi"
	webapi "git.ezbuy.me/ezbuy/evtalk/rpc/proto/etwebapi"

	"fmt"

	"github.com/gin-gonic/gin"
)

type Hello struct {
}

func (s *Hello) SayHello(c *gin.Context, req *webapi.HelloRequest) (*webapi.HelloReply, error) {
	return &webapi.HelloReply{
		Message: "response",
	}, nil
}

func (s *Hello) GetPictureList(c *gin.Context, req *webapi.GetPictureListReq) (*webapi.GetPictureListResp, error) {
	fmt.Println("====>>>helllo")
	resp := &webapi.GetPictureListResp{
		StatusCode: 200,
	}
	pictures := []*webapi.Picture{
		{Id: "fengjing05",
			Title:  "风景",
			Name:   "fengjing05",
			ImgNum: 10,
			ImgSrc: "https://img-cdn-qiniu.dcloud.net.cn/tuku/img/fengjing05.jpg"},
		{Id: "qiche10",
			Title:  "汽车",
			Name:   "qiche10",
			ImgNum: 10,
			ImgSrc: "https://img-cdn-qiniu.dcloud.net.cn/tuku/img/qiche10.jpg"},
		{Id: "fengjing02",
			Title:  "风景",
			Name:   "fengjing02",
			ImgNum: 10,
			ImgSrc: "https://img-cdn-qiniu.dcloud.net.cn/tuku/img/fengjing02.jpg"},
		{Id: "yundong06",
			Title:  "运动",
			Name:   "yundong06",
			ImgNum: 10,
			ImgSrc: "https://img-cdn-qiniu.dcloud.net.cn/tuku/img/yundong06.jpg"},
		{Id: "meinv06",
			Title:  "美女",
			Name:   "meinv06",
			ImgNum: 10,
			ImgSrc: "https://img-cdn-qiniu.dcloud.net.cn/tuku/img/meinv06.jpg"},
	}

	resp.Data = pictures
	return resp, nil
}
