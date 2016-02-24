package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"zmy/model"
)

var db gorm.DB
var db_err error

func init() {
	db, db_err = gorm.Open("sqlite3", "zmy.db")
	if db_err != nil {
		fmt.Println(db_err.Error())
	}
	db.AutoMigrate(&model.Point{})
}

/*Arduino 上传数据类实体*/

func D_Locate(c *gin.Context) {
	fmt.Println("get it!!!!!!!!!")

	fmt.Println(c.Request.Body)
	fmt.Println(c.Request.FormValue("lng"))
	fmt.Println(c.Request.FormValue("lat"))

}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	r.POST("/device/locate", D_Locate)

	/*控制台路由*/
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	authorized.GET("/", Index)

	r.Run(":8080")
}

/**********************************
         控制台路由方法
**********************************/
func Index(c *gin.Context) {
	point := model.Point{}
	point.GetLast(db)

	obj := make(map[string]interface{})
	obj["lp"] = point

	c.HTML(http.StatusOK, "index.html", obj)
}
