import axios from 'axios';
import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';

export const Register = () => {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');


    const handleRegister = async () => {
        try {
            const response = await axios.post('http://localhost:8080/api/register', {
                username: username,
                password: password
            });

            const successMessage = response.data;
            console.log(successMessage);
            setErrorMessage('');
            hanldeRegisterSuccessfulClick();
            window.alert('Registration successful!');
        } catch (err) {
            console.error('Registration failed:', errorMessage.response ? errorMessage.response.data : errorMessage.message);
            setErrorMessage('Registration failed. Please check your credentials.');
        }
    }

    const navigate = useNavigate();
    const hanldeRegisterSuccessfulClick = () => {
        navigate('/');
    }

    return (
        <div>
            <h1>Register</h1>
            <form>
                <label>Username:</label><br />
                <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} /><br />
                <label>Password:</label><br />
                <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} /><br />
                <button onClick={(e) => {
                    e.preventDefault();
                    handleRegister();
                }}>Register</button>
            </form>
        </div >
    );
};