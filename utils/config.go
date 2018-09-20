package utils

import (
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
)

//定义config变量
var(
	G_server_addr string  //服务器ip地址
	G_server_port string //服务器端口
	G_redis_addr string  //redis ip地址
	G_redis_port string //redis port端口
	G_redis_dbnum string //redis db编号
	G_mysql_addr string //mysql ip地址
	G_mysql_port string //mysql port端品
	G_mysql_dbname string //mysql db库名
	G_mysql_username string //mysql 用户名
	G_mysql_password string //mysql 密码
	G_fdfs_http_addr string //fdfs nginx ip地址
)

func InitConfig()  {
	//从配置文件读取配置信息
	appconf,err:=config.NewConfig("ini","conf/app.conf")
	if err!=nil{
		beego.Debug(err)
		return
	}
	G_server_addr=appconf.String("httpaddr")
	G_server_port=appconf.String("httpport")
	G_redis_addr=appconf.String("redisaddr")
	G_redis_port=appconf.String("redisport")
	G_redis_dbnum=appconf.String("redisdbnum")
	G_mysql_addr=appconf.String("mysqladdr")
	G_mysql_port=appconf.String("mysqlport")
	G_mysql_dbname=appconf.String("mysqldbname")
	G_mysql_username=appconf.String("mysqlusername")
	G_mysql_password=appconf.String("mysqlpassword")
	G_fdfs_http_addr=appconf.String("fdfs_http_addr")
}
func init()  {
	InitConfig()
}