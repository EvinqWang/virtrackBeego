package main

import (
	"github.com/astaxie/beego"
	_ "gobeetestpro/routers"
	"gobeetestpro/utils/DButils"
)

func init() {
	DButils.MysqlClient()
	//utils.RedisClient()
}

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true             //开启Session模块
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400 //设置Session有效期,单位秒

	beego.Run()
}
