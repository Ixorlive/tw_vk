import React from 'react';
import LoginForm from '../components/Auth/LoginForm';


const Login: React.FC = () => {
    return (
        <div className="login-page" >
            <h1>Login </h1>
            < LoginForm />
        </div>
    );
}

export default Login;
