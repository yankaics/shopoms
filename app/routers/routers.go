package routers

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/shopoms/app/controllers"
)

func init() {
	// 路由设置
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/bindweixin", &controllers.MainController{}, "*:BindWeixin")
	beego.Router("/setting", &controllers.MainController{}, "*:Setting")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Router("/help", &controllers.HelpController{}, "*:Index")
	beego.Router("/install", &controllers.InstallController{}, "*:Index")
	beego.Router("/upgrade", &controllers.UpgradeController{}, "*:Index")
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.AgentController{})

}
