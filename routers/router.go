package routers

import (
	"turtleOverlordServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
		beego.Router("/bootstrap", &controllers.BootstrapController{})
}
