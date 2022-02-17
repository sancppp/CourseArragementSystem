package main

import (
	"CAS/db/mysql"
	"CAS/db/redis"
	"CAS/model"
	"CAS/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	mysql.Default()          //初始化数据库链接
	model.Default()          //初始化gorm
	redis.Client()           //初始化redis数据库中数据，方便debug
	router.RegisterRouter(r) //注册路由
	err := r.Run(":80")
	if err != nil {
		return
	} //listen and serve on 0.0.0.0:80
}
