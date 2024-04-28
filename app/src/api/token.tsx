import { AUTH_API_TOKEN } from "../config"

export async function validateToken(token: string) {
    const endpoint = AUTH_API_TOKEN;
    try {
        const response = await fetch(endpoint, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ token })
        });
        if (!response.ok) throw new Error('Token validation failed');
        return await response.json();
    } catch (error) {
        console.error('Error validating token:', error);
        throw error;
    }
}