import React, { useState, useEffect } from 'react'
import "./HomePage.css"
import Nav from "../../components/Nav/Nav"
import Input from '../../components/Input/Input'
import CreateTodoButton from '../../components/CreateTodoButton/CreateTodoButton'
import TodoForm from '../../components/TodoForm/TodoForm'
import uuid from 'react-uuid'
import { useHistory } from 'react-router';
import auth from '../../auth/auth'

export default function HomePage(props) {
    const [todoName, setTodoName] = useState("")
    const [username, setUsername] = useState(props.location.state.username)
    const [todos, setTodo] = useState([])
    let history = useHistory()

    useEffect(() => {
        document.body.className = "home-body-color"
        checkAuth()
        getTodos()
        console.log("props", username);
    }, []);

    const checkAuth = () => {
        if(!auth.authenticated){
            history.push("/login")
        }
    }

    const getTodos = async () => {
        const response = await fetch(`http://localhost:8080/api/v1/todos/failson`, {
            method: "GET",
            headers: {
                "Content-type": "application/json"
            }
        })
        const result = await response.json()

        console.log("result", result)
    }

    const editTodo = (id, newTodo) => {
        //If the user passed in no todo, simply end the function call.
        if(newTodo.trim() === ""){
            return
        }

        const todosCopy = [...todos]

        todosCopy.forEach(elem => {
            if(elem.id === id){
                elem.todoName = newTodo
            }
        })
       
        setTodo(todosCopy)
    }

    const removeTodo = (id) => {
        const updatedTodos = todos.filter(todo => id !== todo.id)
        setTodo(updatedTodos)
    }

    const addNewTodo = e => {
        if(todoName.trim() === ""){
            return
        }

        setTodoName("")
        setTodo([...todos, {id: uuid(), todoName: todoName, completed: false}])       
    }

    
    return (
        <div>
            <Nav/>
            <div className="create-todo">
                <Input color={"white"} value={todoName} onChange={e => setTodoName(e.target.value)} placeholder={"Enter To-do"}/>
                <CreateTodoButton onClick={addNewTodo} buttonName={"Add Task"}/>
            </div>               
            <TodoForm todos={todos} removeTodo={removeTodo} updateTodo={editTodo}/>
        </div>
    )
}
