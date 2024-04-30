import React from 'react';
import AuthForm from '../components/Auth/AuthForm';


const Login: React.FC = () => {
    return (
        <div className="login-page" >
            {/* <h1>Login </h1> */}
            < AuthForm isRegister={false} />
        </div>
    );
}

export default Login;
