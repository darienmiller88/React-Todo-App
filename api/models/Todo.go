package models

import (

	validation "github.com/go-ozzo/ozzo-validation"
)

type Todo struct {
	//mgm.DefaultModel `bson:",inline,omitempty"`
	TodoTask      string `json:"todo_task" bson:"todo_task"`
	Date          string `json:"date"      bson:"date"`
	TodoCompleted bool   `json:"completed" bson:"completed"`
	ID            int    `json:"id"        bson:"id"` 
	//	Username string  `json:"username"  bson:"username"`
}

func (todo Todo) Validate() error {
	return validation.ValidateStruct(&todo,

		//todo_task is required, and must be between 3 and 45 characters
		validation.Field(&todo.TodoTask, validation.Required, validation.Length(3, 45)),
	//	validation.Field(todo.Username, validation.Required, validation.By(todo.checkUsernameExists)),
	)
}

// func (todo *Todo) checkUsernameExists(value interface{}) error {
// 	username, _ := value.(string)
// 	result := &Todo{}

// 	mgm.Coll(&Todo{}).FindOne(mgm.Ctx(), bson.M{"username": username}).Decode(result)
// 	fmt.Println(result)
// 	if result.Username != "" {
// 		return errors.New(fmt.Sprintf("Username %s doesn't exist!", username))
// 	}

// 	return nil
// }
