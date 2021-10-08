import React from 'react'
import { Lock, PersonCircle } from 'react-bootstrap-icons'
import { Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import "../../colors/colors.css"
import "./PostForm.css"


export default function PostForm(props) {
    return (
        <form className="post-form" onSubmit={props.onSubmit}>
            <div className="header">{props.header}</div>   
            <h5>And enjoy your todo list!</h5>        
            <div className="input-container">
                <PersonCircle className="person-circle"/>
                <input type="text" className="input-bar" required value={props.username} placeholder="Username" onChange={e => props.setUsername(e.target.value)}/> 
            </div>
            
            {
                props.usernameError 
                ? 
                <li style={{color: "red"}}>{props.usernameError}<br/><br/></li>
                :
                <div/>
            }
    
            <div className="input-container">
                <Lock className="lock"/>
                <input type="password" className="input-bar" required value={props.password} placeholder="Password" onChange={e => props.setPassword(e.target.value)}/>
            </div>
            <br/>
            
            <div>
                {
                    props.appendErrors()
                }
            </div>           
            
            <Button type="submit" className="submit-form-btn">{props.buttonName}</Button>
            <br/>
            <br/>
            <div className="sign-in">
                {props.redirectMessage} <Link to={props.route}>{props.linkname}</Link>   
            </div>
            <br/>
        </form>       
    )
}
