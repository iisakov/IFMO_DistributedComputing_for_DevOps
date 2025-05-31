package main

import (
	"app/config"
	_ "app/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           app by artisan
// @version         0.1
// @description     app для распределённых систем.
// @contact.name   Artisan
// @contact.url    http://www.by_artisan.io/support
// @contact.email  by@artisan.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      127.0.0.1:8080
// @BasePath  /
func main() {
	r := config.InitEngine()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
