import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css'
import HomePage from './pages/HomePage/HomePage'
import LoginPage from './pages/LoginPage/LoginPage'
import SignupPage from './pages/SignupPage/SignupPage'
import  {BrowserRouter as Router, Switch, Route, Redirect} from 'react-router-dom'

function App() {
    return (
      <Router>
        <Switch>
          <Route exact path="/" render={()=>{return <Redirect to="/login"/>}}/>
          <Route exact path="/home" component={HomePage}/>
          <Route exact path="/signup" component={SignupPage}/>
          <Route exact path="/login" component={LoginPage}/>
          <Route path="*" component={() => {return <div>404 not found</div>}}/>
        </Switch>
      </Router>
    )
}

export default App;
