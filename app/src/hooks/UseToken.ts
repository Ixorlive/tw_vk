import { useState, useEffect } from 'react';
import { validateToken } from '../api/token';
import { UserType } from "../entity/UserType"


const initialUser = { id: 0, login: '' };

function useToken() {
    const [user, setUser] = useState<UserType>(initialUser);

    const getToken = () => {
        return localStorage.getItem('token');
    };

    const clearToken = () => {
        localStorage.removeItem('token');
    };

    useEffect(() => {
        const token = getToken();
        if (!token) {
            setUser(initialUser);
        } else {
            validateToken(token).then(
                (data) => {
                    if (data.id && data.login) {
                        setUser({ id: data.id, login: data.login });
                    } else {
                        clearToken();
                        setUser(initialUser);  // Token invalid, reset user
                    }
                },
                (error) => {
                    console.error('Token validation error:', error);
                    clearToken();
                    setUser(initialUser);  // On error, reset user
                }
            );
        }
    }, []);

    return user;
}

export default useToken;