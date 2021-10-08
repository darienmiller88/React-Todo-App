package controllers

import(
	"github.com/gin-gonic/gin"
)

type ViewController struct{
	Router *gin.RouterGroup
}

func (v *ViewController) Init(routeToMount *gin.RouterGroup){
	v.Router = routeToMount

	v.Router.Static("/static", "../client/static")
	v.Router.GET("/", v.home)
}

func (v *ViewController) home(c *gin.Context){
	c.HTML(200, "index", nil)
}