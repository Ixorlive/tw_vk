// src/components/Auth/LoginForm.tsx
import React, { useState } from 'react';
import { AUTH_API_LOGIN, AUTH_API_REGISTER } from '../../config';
import { useNavigate } from 'react-router-dom';
import './AuthForm.css'

interface AuthFormProps {
    isRegister: boolean
}

const AuthForm: React.FC<AuthFormProps> = ({ isRegister }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [status, setStatus] = useState('');
    const navigate = useNavigate();
    const url = isRegister ? AUTH_API_REGISTER : AUTH_API_LOGIN;
    const actionStr = isRegister ? "Register" : "Login";

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        try {
            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    login: username,
                    password: password
                })
            });
            const data = await response.json();
            if (!response.ok) {
                throw new Error(`Failed to ${actionStr}: ` + data.error);
            }
            if (isRegister) {
                console.log("Registration in successfully");
                navigate("/login")
            } else if (data.token) {
                localStorage.setItem('token', data.token);
                setStatus("Logged in successfully")
                window.location.reload()
            } else {
                throw new Error('No token received');
            }

        } catch (error) {
            setStatus(`${actionStr} error: ${error}`)
        }
    };

    return (
        <div className="login-container">
            <h1>{actionStr}</h1>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="username">Username:</label>
                    <input
                        type="text"
                        id="username"
                        name="username"
                        required
                        value={username}
                        onChange={e => setUsername(e.target.value)}
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        required
                        value={password}
                        onChange={e => setPassword(e.target.value)}
                    />
                </div>
                <p>{status}</p>
                <button type="submit" className="log_btn">{actionStr}</button>
            </form>
        </div>
    );
}
export default AuthForm;
