package controllers

import (
	"fmt"
	"TodoApp/api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type TodoController struct {
}

func (u *TodoController) Init(routeToMount *gin.RouterGroup) {
	routeToMount.GET("/:username", u.getTodos)
	routeToMount.POST("/:username", u.addTodo)
	routeToMount.DELETE("/:username/:id", u.deleteTodo)	
	routeToMount.PUT("/:username/:id", u.editTodo)	
	routeToMount.GET("/:username/:id", u.getTodoById)
}

//CREATE - Add a todo to the database
func (u *TodoController) addTodo(c *gin.Context) {
	todo := models.Todo{}
	user := models.User{}
	username := c.Param("username")

	if err := c.Bind(&todo); err != nil {
		c.JSON(c.Writer.Status(), err.Error())
		return
	}

	if err := todo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result := mgm.Coll(&models.User{}).FindOne(mgm.Ctx(), bson.M{"username": username})

	if result.Err() != nil{
		c.JSON(http.StatusNotFound, result.Err().Error())
		return
	}

	result.Decode(&user)
	todo.Date = time.Now().Local().String()
	todo.ID = user.Todos[len(user.Todos)-1].ID + 1

	update, err := mgm.Coll(&models.User{}).UpdateOne(
		mgm.Ctx(),
		bson.M{"username": username},
		bson.M{"$push": bson.M{"todos": todo}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, update)
}

//READ - Get a todo by its ID
func (u *TodoController) getTodoById(c *gin.Context) {
	user := models.User{}
	username := c.Param("username")
	id, _ := strconv.Atoi(c.Param("id"))
	result := mgm.Coll(&models.User{}).FindOne(mgm.Ctx(), bson.M{"username": username})

	if result.Err() != nil {
		c.JSON(http.StatusBadRequest, result.Err().Error())
		return
	}

	if err := result.Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, todo := range user.Todos {
		if id == todo.ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"err": fmt.Sprintf("No todo with ID %d not found", id)})
}

//READ - Get ALL todos by username
func (u *TodoController) getTodos(c *gin.Context) {
	username := c.Param("username")
	user := models.User{}
	result := mgm.Coll(&models.User{}).FindOne(mgm.Ctx(), bson.M{"username": username})

	if result.Err() != nil {
		c.JSON(http.StatusInternalServerError, result.Err().Error())
		return
	}

	if err := result.Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, user.Todos)
}

//UPDATE - Edit a todo
func (u *TodoController) editTodo(c *gin.Context) {
	todo := models.Todo{}
	id, _ :=strconv.Atoi(c.Param("id"))
	username := c.Param("username")

	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := todo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := mgm.Coll(&models.User{}).UpdateOne(
		mgm.Ctx(),

		//Find the collection with the username passed, and the id of the todo to be updated.
		bson.M{"username": username, "todos.id": id},

		//Overide the todo task of the todo with a new one passed by the user
		bson.M{"$set": bson.M{"todos.$.todo_task": todo.TodoTask}},
	)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

//DELETE - Delete a todo
func (u *TodoController) deleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	username := c.Param("username")
	result, err := mgm.Coll(&models.User{}).UpdateOne(
		mgm.Ctx(),

		bson.M{"username": username},

		//Remove an todo from the todos array of a user with an id "id".
		bson.M{"$pull": bson.M{"todos": bson.M{"id": id}}},
	)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
