package main

import (
	db "home-work-z-api/database"
	"home-work-z-api/route"
)

// @title Swagger Home Work Z API
// @version 1.0
// @description Home work Z Swagger 說明文件
// @termsOfService http://swagger.io/terms/
// @contact.name Alan Syu
// @contact.email qoo753951ooq@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8080
// @BasePath /home-work-z-api
// @schemes http
func main() {
	db.InitDB()
	route.InitRouter()
}
