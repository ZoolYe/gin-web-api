package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

var engine *xorm.Engine

func init() {

	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/go?charset=utf8")
	if err != nil {
		log.Println("数据库引擎创建失败", err.Error())
		return
	}

	if err := engine.Ping(); err != nil {
		log.Println("数据库链接失败", err.Error())
		return
	}

	//打印SQL日志
	engine.ShowSQL()
	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)
	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	engine.SetTableMapper(core.SnakeMapper{})
}

type userinfo struct {
	Uid        int       `json:"uid"`
	Username   string    `json:"userName"`
	Department string    `json:"department"`
	Created    time.Time `json:"created"`
}

func main() {

	r := gin.Default()

	r.GET("/getUserById", getUser)
	r.POST("/addUser", addUser)
	r.POST("/deleteUser", deleteUser)
	r.POST("/updateUser", updateUser)

	r.Run(":6240")
}

func getUser(c *gin.Context) {

	var user userinfo

	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		log.Println(err.Error())
		return
	}
	user.Uid = uid
	engine.Get(&user)
	c.JSON(http.StatusOK, user)
}

func addUser(c *gin.Context) {
	var user userinfo
	user.Created = time.Now()
	c.BindJSON(&user)
	engine.Insert(&user)
	c.JSON(http.StatusOK, gin.H{"msg": "添加完成", "status": http.StatusOK})
}

func deleteUser(c *gin.Context) {
	var user userinfo
	c.BindJSON(&user)
	engine.Delete(user)
	c.JSON(http.StatusOK, gin.H{"msg": "删除完成", "status": http.StatusOK})
}

func updateUser(c *gin.Context) {
	var user userinfo
	c.BindJSON(&user)
	engine.Where("uid = ?", user.Uid).Update(user)

	c.JSON(http.StatusOK, gin.H{"msg": "更新成功", "status": http.StatusOK})
}
