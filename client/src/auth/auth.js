const auth = {
    authenticated: false,
    errors: {},
    async signin(username, password) {
        const data = {
            username,
            password
        }

        console.log(data)

        const response = await fetch("http://localhost:8080/api/v1/users/signin", {
            method: "POST", 
            headers: {
                "Content-type": "application/json"
            },
            body: JSON.stringify(data)
        })

        const errors = await response.json()

        if(!errors["error"]){
            this.authenticated = true
        }

        this.errors = errors
    },
    async signout(history){
        const response = await fetch("http://localhost:8080/api/v1/users/signout", {
            method: "POST",
            headers: {
                "Content-type": "application/json"
            }
        })
        const result = await response.json()
        history.push("/login")
        this.authenticated = false
    }
}

export default auth