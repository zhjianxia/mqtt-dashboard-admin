package main

import (
	ginAdapter "github.com/GoAdminGroup/go-admin/adapter/gin"     // 适配器
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql 驱动

	_ "github.com/GoAdminGroup/themes/adminlte" // ui主题
	//_ "github.com/GoAdminGroup/themes/sword" // ui主题sword

	"io/ioutil"

	"tables"

	//"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"

	//"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// 实例化一个GoAdmin引擎对象
	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(config.Config{
		// 数据库配置，为一个map，key为连接名，value为对应连接信息
		Databases: config.DatabaseList{
			// 默认数据库连接，名字必须为default
			"default": {
				Host:       "127.0.0.1",
				Port:       "13306",
				User:       "root",
				Pwd:        "123456",
				Name:       "cloudmq",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     db.DriverMysql,
			},
			// 默认数据库连接，名字必须为default
			"blogdb": {
				Host:       "127.0.0.1",
				Port:       "13306",
				User:       "root",
				Pwd:        "123456",
				Name:       "blog",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     db.DriverMysql,
			},
		},
		UrlPrefix: "admin", // 访问网站的前缀
		IndexUrl:  "/",
		// Store 必须设置且保证有写权限，否则增加不了新的管理员用户
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Debug:    true,
		Language: language.CN,
		//Theme:     "sword", //配置主题名字
	}).
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	//eng.HTML("GET", "/admin", datamodel.GetContent)

	//r.GET("/admin", ada.Content(func(ctx *ada.Context) (panel types.Panel, e error) {
	//	return GetDashboardContent(ctx)
	//}))

	eng.HTML("GET", "/admin", GetDashboardContent)
	r.GET("/admin/custom", ginAdapter.Content(GetMyContent))

	_ = r.Run(":9033")
}
