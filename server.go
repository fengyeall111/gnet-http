package gnet_http

import (
	"time"

	"github.com/panjf2000/gnet/v2"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	*fasthttp.Server
	gnet.Engine
}

func NewHttpServer(srv *fasthttp.Server) *HttpServer {
	return &HttpServer{
		Server: srv,
	}
}

func (wss *HttpServer) OnBoot(eng gnet.Engine) gnet.Action {
	wss.Engine = eng
	return gnet.None
}

func (wss *HttpServer) OnOpen(c gnet.Conn) ([]byte, gnet.Action) {
	if err := wss.Server.ServeConn(c); err != nil {
		return nil, gnet.Close
	}
	return nil, gnet.None
}

func (wss *HttpServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	return gnet.None
}

func (wss *HttpServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	return gnet.None
}

func (wss *HttpServer) OnTick() (delay time.Duration, action gnet.Action) {
	return 3 * time.Second, gnet.None
}

func (wss *HttpServer) OnShutdown(gnet.Engine) {
	return
}

func (h *HttpServer) Run(addr string, opts ...gnet.Option) error {
	return gnet.Run(h, addr, opts...)
}
