import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Login from './pages/Login';
import Register from './pages/Register';
import Main from './pages/Main';
import useToken from './hooks/UseToken';

function App() {
    const user = useToken()

    return (
        <BrowserRouter>
            <Routes>
                <Route
                    path="/login"
                    element={user.id === 0 ? <Login /> : <Navigate to="/" />} />
                <Route
                    path="/register"
                    element={user.id === 0 ? <Register /> : <Navigate to="/" />} />
                <Route
                    path="/"
                    element={<Main user={user} />}
                />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
