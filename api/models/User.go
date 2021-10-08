package models

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Username         string `json:"username" bson:"username"`
	Password         string `json:"password" bson:"password"`
	Todos            []Todo `json:"todos"    bson:"todos"`
	Authenticated    bool
}

func (user *User) ValidateUserCredentials() gin.H {
	response := make(gin.H)
	usernameErr := validation.Validate(user.Username, validation.Required, validation.Length(5, 30),
		validation.By(user.checkUsernameExists))
	passwordErrs := user.validatePassword()

	if usernameErr != nil {
		response["username"] = usernameErr.Error()
	}

	if len(passwordErrs) > 0 {
		response["password"] = passwordErrs
	}

	return response
}

func (user *User) validatePassword() []gin.H {
	errors := []gin.H{}
	passwordLen := 7
	passwordLenErr := validation.Validate(user.Password, validation.Required, validation.Length(passwordLen, 0))
	digitErr := validation.Match(regexp.MustCompile("[0-9]")).Validate(user.Password)
	lowercaseErr := validation.Match(regexp.MustCompile("[a-z]")).Validate(user.Password)
	uppcaseErr := validation.Match(regexp.MustCompile("[A-Z]")).Validate(user.Password)

	if passwordLenErr != nil || digitErr != nil || lowercaseErr != nil || uppcaseErr != nil {
		errors = []gin.H{
			{
				"password_err":      fmt.Sprintf("Password must be at least %d characters.", passwordLen),
				"validation_passed": passwordLenErr == nil,
			},
			{
				"password_err":      "Password must contain at least one digit.",
				"validation_passed": digitErr == nil,
			},
			{
				"password_err":      "Password must contain at least one lowercase letter.",
				"validation_passed": lowercaseErr == nil,
			},
			{
				"password_err":      "Password must contain at least one uppercase letter.",
				"validation_passed": uppcaseErr == nil,
			},
		}
	}

	return errors
}

func (user *User) checkUsernameExists(value interface{}) error {
	username, _ := value.(string)
	result := &User{}

	mgm.Coll(&User{}).FindOne(mgm.Ctx(), bson.M{"username": username}).Decode(result)
	fmt.Println(result)
	if result.Username != "" {
		return errors.New(fmt.Sprintf("Username %s is taken! Please enter a different one.", username))
	}

	return nil
}

func (user *User) InitializeTodos() {
	for i := 0; i < len(user.Todos); i++ {
		user.Todos[i].Date = time.Now().String()
		user.Todos[i].ID = i + 1
	}
}
