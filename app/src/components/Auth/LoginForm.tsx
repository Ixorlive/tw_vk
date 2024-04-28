// src/components/Auth/LoginForm.tsx
import React, { useState } from 'react';
import { AUTH_API_LOGIN } from '../../config';
import { useNavigate } from 'react-router-dom';

const LoginForm = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        try {
            const response = await fetch(AUTH_API_LOGIN, {
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
                throw new Error('Failed to login: ' + data.error);
            }

            if (data.token) {
                localStorage.setItem('token', data.token);
                console.log("Logged in successfully");
                navigate("/")
            } else {
                throw new Error('No token received');
            }
        } catch (error) {
            console.error('Login error:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Username:
                <input type="text" value={username} onChange={e => setUsername(e.target.value)} />
            </label>
            <label>
                Password:
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} />
            </label>
            <button type="submit">Log In</button>
        </form>
    );
}

export default LoginForm;
