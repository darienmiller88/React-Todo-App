import React from 'react'
import "../../colors/colors.css"
import "./CreateTodoButton.css"

export default function CreateTodoButton(props) {
    return (
        <button id="add-task-button" onClick={props.onClick}>{props.buttonName}</button>
    )
}
