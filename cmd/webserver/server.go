package webserver

import (
	"1536509937/ku-bbs/pkg/config"
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"1536509937/ku-bbs/internal/route"
	"1536509937/ku-bbs/pkg/utils"
)

func Run() {
	if config.Conf.System.Env == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.SetFuncMap(utils.GetTemplateFuncMap())
	engine.Static("/assets", "../assets")
	engine.LoadHTMLGlob("../views/**/**/*")

	store := cookie.NewStore([]byte(config.Conf.Session.Secret))
	engine.Use(sessions.Sessions(config.Conf.Session.Name, store))

	route.RegisterBackendRoute(engine)
	route.RegisterFrontedRoute(engine)

	if err := engine.Run(fmt.Sprintf(":%s", config.Conf.System.Addr)); err != nil {
		log.Fatalf("server running error: %v", err)
	}
}
