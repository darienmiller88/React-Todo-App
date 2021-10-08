import React, { useEffect, useState } from 'react'
// import '../../colors/colors.css'
import { useHistory } from 'react-router';
import PostForm from '../../components/PostForm/PostForm';
import './SignupPage.css'
import uuid from 'react-uuid'
import auth from '../../auth/auth';

export default function SignupPage() {
    const [passwordErrors, setPasswordsWords] = useState([])
    const [usernameError, setUserNameError] = useState("")
    // const [usernames, setUsernames] = useState(['darienmiller88', 'nisey', 'user1', 'fred'])
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    let history = useHistory()

    useEffect(() => {
        document.body.className = "signup-body-color";
        if(auth.authenticated){
            history.push("/home")
        }
    }, []);

    async function signup(e) {
        e.preventDefault()
        const data = {
            username,
            password,
        }

        console.log(data)
        const response = await fetch("http://localhost:8080/api/v1/users/signup", {
            method: "POST",
            headers: {
                "Content-type": "application/json"
            },
            body: JSON.stringify(data)
        })
        const result = await response.json()
        
        console.log(result)

        if(result["success"]){
            history.push("/home")
        }
        else{
            setUserNameError(result["username"])
            setPasswordsWords(result["password"])      
            setUsername("")
            setPassword("")
        }
    }

    function appendErrors(){
        return (
            <>
                {            
                    passwordErrors.map(passwordError => {
                        if(!passwordError.validation_passed){
                            return <li key={uuid()} style={{color: "red"}}>{passwordError.password_err}</li>      
                        }
                        return <li key={uuid()} style={{color: "green"}}>{passwordError.password_err}</li>
                    })
                }   
                <br/>
            </>          
        )    
    }

    return (
        <PostForm
            header={"Sign up"}
            username={username}
            setUsername={setUsername}
            password={password}
            setPassword={setPassword}
            onSubmit={signup}
            usernameError={usernameError}
            passwordErrors={passwordErrors}
            appendErrors={appendErrors}
            buttonName={"Create Account"}
            redirectMessage={"Already have an account?"}
            linkname={"Sign in"}
            route={"/login"}
        />
        // <form className="signup-form" onSubmit={signup}>
        //     <div className="header">Sign up</div>   
        //     <i>And enjoy your todo list!</i>        
        //     <div className="input-container">
        //         <PersonCircle  className="person-circle"/>
        //         <input type="text" required value={username} placeholder="Username" onChange={e => setUsername(e.target.value)}/> 
        //     </div>

        //     {
        //         usernameTaken 
        //         ? <div className="username-taken" style={{color: "red", marginBottom:"20px"}}>
        //             Username "darienmiller88" is taken!
        //         </div> : <div></div>
        //     }
            
        //     <div className="input-container">
        //         <Lock className="lock"/>
        //         <input type="password" required value={password} placeholder="Password" onChange={e => setPassword(e.target.value)}/>
        //     </div>
        //     <br/>
        //     <Button type="submit" className="create-account">Create account</Button>
        //     <br/>
        //     <br/>
        //     <div className="sign-up">
        //         Already have an account? <Link to="/login"> Sign in</Link>   
        //     </div>
        //     <br/>
        // </form>
    )
}
