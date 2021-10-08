import React, {  useState, useEffect } from 'react'
import { useHistory } from 'react-router-dom';
import PostForm from '../../components/PostForm/PostForm';
import auth from "../../auth/auth"
import "./LoginPage.css"

export default function LoginPage() {
    const [loginError, setLoginError] = useState("")
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    let history = useHistory()

    async function login(e) {
        e.preventDefault()

       auth.signin()
        
        if(!auth.errors["error"]){            
            //setLoginError("")
            history.push("/home", {username,})       
        }else{
            setLoginError(!auth.errors["error"] ? "" : auth.errors["error"])
            setUsername("")
            setPassword("")    
        }   
    }

    useEffect(() => {
        document.body.className = "login-body-color";
        if(auth.authenticated){
            history.push("/home")
        }
    }, []);

    function appendErrors(){
        return (
            <>
                {                
                    loginError !== "" ? <li style={{color: "red"}}>{loginError}</li> : null
                }
                <br/>
            </>
        )
    }

    return (
        <PostForm
            header={"Sign in"}
            username={username}
            setUsername={setUsername}
            password={password}
            setPassword={setPassword}
            appendErrors={appendErrors}
            onSubmit={login}
            buttonName={"Log in"}
            redirectMessage={"Don't have an account?"}
            linkname={"Sign up"}
            route={"/signup"}
        />
    )
}
