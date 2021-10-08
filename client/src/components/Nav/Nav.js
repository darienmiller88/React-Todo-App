import React from 'react'
import "../../colors/colors.css"
import "./Nav.css"
import { useHistory } from 'react-router';
import auth from "../../auth/auth"

export default function Nav() {
    let history = useHistory()

//    async function logoutHandler(){
//         const response = await fetch("http://localhost:8080/api/v1/users/signout", {
//             method: "POST",
//             headers: {
//                 "Content-type": "application/json"
//             }
//         })
//         const result = await response.json()

//         console.log(result);
//         history.push("/login")
//     }

    return (
        <div className="mini-nav">
            <div className="app-name">Todo App</div>
            <div className="logout" onClick={auth.signout(history)}>Logout</div>
        </div> 
    )
}
