package controllers

import(
	"github.com/gin-gonic/gin"
)

type MountController struct{

}

func (m *MountController) Init(routerToMount *gin.RouterGroup){
	userController := UserController{}
	todoController := TodoController{}

	userController.Init(routerToMount.Group("/users"))
	todoController.Init(routerToMount.Group("/todos"))
}