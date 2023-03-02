package api

import (
	"github.com/fasthttp/router"
	config "github.com/gusleein/goconfig"
	log "github.com/gusleein/golog"
	"github.com/valyala/fasthttp"
)

var server *fasthttp.Server

func Start() {
	log.Info("http server start...")
	server = &fasthttp.Server{
		Handler: createRequestHandler(),
	}
	if err := server.ListenAndServe(config.GetString("port")); err != nil {
		log.Fatal(err.Error())
	}
}
func Stop() {
	if err := server.Shutdown(); err != nil {
		log.Error(err.Error())
	}
	log.Info("http server stop...")
}

func createRequestHandler() func(ctx *fasthttp.RequestCtx) {
	r := router.New()
	r.POST("/api/v1/", createUserInfo)
	r.GET("/api/v1/{userId}", getInfoUserById)
	r.GET("/api/v1/top-{field:[a-zA-Z-]+}", GetTop100)
	return r.Handler
}
