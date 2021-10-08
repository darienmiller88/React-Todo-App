import React, { createRef, useState } from 'react'
import "../../colors/colors.css"
import "./TodoBody.css"
import Input from '../Input/Input';
import { Pencil, Trash } from 'react-bootstrap-icons';
import { Modal, Button } from 'react-bootstrap'
import CreateTodoButton from '../CreateTodoButton/CreateTodoButton'

function VerticallyCenteredModal(props) {
    return (
      <Modal {...props} size="lg" aria-labelledby="contained-modal-title-vcenter" centered>
        <Modal.Header>
          <Modal.Title id="contained-modal-title-vcenter">
            {
                props.isedit ? <div>Edit your todo!</div> : <div className="todo-title">Delete todo?</div>
            }
          </Modal.Title>
        </Modal.Header>
        {
            props.isedit
            ? 
            <Modal.Body>
                <div className="create-todo">
                    <Input color={"black"} onChange={props.onChange} placeholder={"Enter new To-do"}/> 
                    <CreateTodoButton buttonName={"Edit"} onClick={props.updateTodo}/>
                </div>           
            </Modal.Body>
            : null
        }
        
        <Modal.Footer>
            {
                !props.isedit ? <Button onClick={props.removeTodo}>Delete Todo</Button> : null
            }
            <Button onClick={props.onHide}>Close</Button>
        </Modal.Footer>
      </Modal>
    );
  }

export default function TodoBody(props) {
    const [modalShow, setModalShow] = useState(false);
    const [isEdit, setEdit] = useState(false)
    const [taskName, setTaskName] = useState()
    const [checked, setChecked] = useState()
    const ref = createRef(null)

    const editAndClose = (id, todoName) => {
        props.updateTodo(id, todoName)
        setModalShow(false)
    }

    const deleteAndClose = (id) => {
        props.removeTodo(id)
        setModalShow(false)
    }

    const setToEditOrDelete = (isEdit) => {
        setModalShow(true)
        setEdit(isEdit)
    }

    return (
        <div>
            <div className="todo">
                <input 
                    type="checkbox" 
                    onChange={() => setChecked(!checked)} 
                    id="todo-completed"
                />
                <div className={checked ? "todo-name todo-checked" : "todo-name"} >{props.todo.todoName}</div>
                <button  className="base-button edit-button"  onClick={() => setToEditOrDelete(true)}><Pencil/></button>
                <button className="base-button delete-button" onClick={() => setToEditOrDelete(false)}><Trash/></button>
            </div>  
            <VerticallyCenteredModal 
                value={taskName}
                onChange={e => setTaskName(e.target.value)} 
                noderef={ref} 
                show={modalShow} 
                isedit={isEdit ? 1 : 0} 
                onHide={() => setModalShow(false)}
                removeTodo={() => deleteAndClose(props.todo.id)}
                updateTodo={() => editAndClose(props.todo.id, taskName)}
                id={props.todo.id}
            />
        </div>
    )
}