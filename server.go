package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
	"strings"
	"time"
	"zmy/model"
	"zmy/util"
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

/*Arduino 接入方法*/
func D_Locate(c *gin.Context) {
	fmt.Println("get it!!!!!!!!!")

	lng_raw := c.Request.FormValue("lng")
	fmt.Println(lng_raw)
	lng_dot := strings.Index(lng_raw, ".")
	lng_s_str := lng_raw[lng_dot+1:][0:2] + "." + lng_raw[lng_dot+1:][2:]
	lng_s, _ := strconv.ParseFloat(lng_s_str, 64)
	fmt.Println("second: ", lng_s)
	lng_m, _ := strconv.ParseFloat(lng_raw[lng_dot-2:lng_dot], 64)
	fmt.Println("minute: ", lng_m)
	lng_d, _ := strconv.ParseFloat(lng_raw[0:lng_dot-2], 64)
	fmt.Println("degree: ", lng_d)

	lat_raw := c.Request.FormValue("lat")
	fmt.Println(lat_raw)
	lat_dot := strings.Index(lat_raw, ".")
	lat_s_str := lat_raw[lat_dot+1:][0:2] + "." + lat_raw[lat_dot+1:][2:]
	lat_s, _ := strconv.ParseFloat(lat_s_str, 64)
	fmt.Println("second: ", lat_s)
	lat_m, _ := strconv.ParseFloat(lat_raw[lat_dot-2:lat_dot], 64)
	fmt.Println("minute: ", lat_m)
	lat_d, _ := strconv.ParseFloat(lat_raw[0:lat_dot-2], 64)
	fmt.Println("degree: ", lat_d)

	lng := lng_d + lng_m/60 + lng_s/3600
	lat := lat_d + lat_m/60 + lat_s/3600

	fmt.Println("lng  and  lat is : ", lat, lng)

	GD_lat, GD_lng := util.WGStoGCJ(lat, lng)
	fmt.Println(GD_lng, GD_lat)

	point := model.Point{Lng: util.LimitFloat(GD_lng), Lat: util.LimitFloat(GD_lat), CreateAt: time.Now().Format("2006-01-02 15:04:05")}
	point.Insert(db)

}

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	/*Arduino 接入路由*/
	r.POST("/device/locate", D_Locate)

	/*AJAX 请求路由*/
	r.GET("/ajax/location", GetLocation)
	r.GET("/ajax/location/all", GetAllLocation)

	/*控制台路由*/
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	authorized.GET("/", Index)
	authorized.GET("/location/list", List)
	authorized.GET("/location/delete/:id", Delete)

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

func List(c *gin.Context) {
	point := model.Point{}
	points := point.GetAll(db)

	obj := make(map[string]interface{})
	obj["points"] = points

	c.HTML(http.StatusOK, "list.html", obj)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	point := model.Point{}
	point.Get(db, id)
	point.Delete(db)
}

/*********************************
		  AJAX 方法
*********************************/

/*获得机要密码箱当前位置*/
func GetLocation(c *gin.Context) {
	point := model.Point{}
	point.GetLast(db)

	c.JSON(http.StatusOK, point)
}

/*获得机要密码箱所有位置*/
func GetAllLocation(c *gin.Context) {
	point := model.Point{}
	points := point.GetAll(db)
	c.JSON(http.StatusOK, points)
}
