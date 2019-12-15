package main

import (
	"flag"
	// "path/filepath"

	"github.com/b3kt/account-srv/config"
	"github.com/b3kt/account-srv/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/golang/glog"
)

func main() {

	addr := flag.String("addr", config.Server.Addr, "Address to listen and serve")
	flag.Parse()

	if config.Server.Mode == gin.ReleaseMode {
		gin.DisableConsoleColor()
	}

	app := gin.Default()

	// app.Static("/images", filepath.Join(config.Server.StaticDir, "img"))
	// app.StaticFile("/favicon.ico", filepath.Join(config.Server.StaticDir, "img/favicon.ico"))
	// app.LoadHTMLGlob(config.Server.ViewDir + "/*")
	// app.MaxMultipartMemory = config.Server.MaxMultipartMemory << 20

	router.Route(app)

	app.Use(cors.Default())

	// Listen and Serve
	app.Run(*addr)
}
