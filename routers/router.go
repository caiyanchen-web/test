package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
