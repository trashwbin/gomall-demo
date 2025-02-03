package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	//h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
	//	c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	//})

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.Data(consts.StatusOK, consts.MIMETextPlain, []byte("Hello, World!"))
	})

	h.Spin()
}
