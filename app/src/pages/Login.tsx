import React from 'react';
import LoginForm from '../components/Auth/LoginForm';
import { AUTH_API_LOGIN } from '../config';


const Login: React.FC = () => {
    console.log(AUTH_API_LOGIN)
    return (
        <div className="login-page" >
            <h1>Login </h1>
            < LoginForm />
        </div>
    );
}

export default Login;
