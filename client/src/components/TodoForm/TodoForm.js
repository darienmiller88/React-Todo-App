import React from 'react'
import "./TodoForm.css"
import TodoBody from '../TodoBody/TodoBody'

export default function TodoForm(props) {
    return (
        <div className="todo-form">
            {
                props.todos.map(todo => <TodoBody key={todo.id} todo={todo} removeTodo={props.removeTodo} updateTodo={props.updateTodo}/>) 
            }
        </div>
    )
}
