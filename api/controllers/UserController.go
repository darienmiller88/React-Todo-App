package controllers

import (
	"fmt"
	"TodoApp/api/models"
	"net/http"

	"TodoApp/api/session"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	fakeUser models.User
}

func (u *UserController) Init(routeToMount *gin.RouterGroup) {
	u.fakeUser = models.User{Username: "darienmiller88", Password: "Ninten5544"}

	routeToMount.POST("/signin", u.signin)
	routeToMount.POST("/signout", u.signout)
	routeToMount.POST("/signup", u.signup)
	routeToMount.GET("/", u.getAllUsers)
}

func (u *UserController) signin(c *gin.Context) {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		c.JSON(c.Writer.Status(), err)
		return
	}

	if user.Username != u.fakeUser.Username || user.Password != u.fakeUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Username or password incorrect. Please try again.",
			"success": false,
		})
		return
	}

	newSession, err := session.Store.Get(c.Request, session.SessionName)

	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Authenticated = true
	newSession.Values["user"] = user
	
	if err = newSession.Save(c.Request, c.Writer); err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (u *UserController) signup(c *gin.Context) {
	user := models.User{}

	//Error handle for failure to bind the json into the user object
	if err := c.Bind(&user); err != nil {
		c.JSON(c.Writer.Status(), err)
		return
	}

	//If the user's data fails to pass the credentials needed for a valid username and password,
	//send the errors back to the user.
	if response := user.ValidateUserCredentials(); len(response) > 0 {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user.InitializeTodos()
	fmt.Println("user", user)
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	newSession, _ := session.Store.Get(c.Request, session.SessionName)

	user.Authenticated = true
	user.Password = string(hash)
	newSession.Values["user"] = user
	err := newSession.Save(c.Request, c.Writer)

	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	mgm.Coll(&models.User{}).Create(&user)
	c.JSON(http.StatusOK, user)
}

func (u *UserController) signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "signing out"})
}
 
func (u *UserController) getAllUsers(c *gin.Context) {
	users := []models.User{}
	documents, err := mgm.Coll(&models.User{}).Find(mgm.Ctx(), bson.D{})

	if err != nil{
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := documents.All(mgm.Ctx(), &users); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Err": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}