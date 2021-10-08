import React from 'react'
import "../../colors/colors.css"
import "./Input.css"


export default function Input(props) {      
    return (
        // <div class="form__group field">
        //     <input 
        //         type="input" 
        //         className={`form__field ${props.color}`}
        //         placeholder={props.placeholder}
        //         onChange={props.onChange} 
        //         value={props.value} 
        //         id='name' 
        //         required 
        //     />
        //     <label for="name" class="form__label">Name</label>
        // </div>
        <input type="text" 
            value={props.value} 
            className={props.color + " todo-input"} 
            onChange={props.onChange} 
            placeholder={props.placeholder}
        />
    )
}
