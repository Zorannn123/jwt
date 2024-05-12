import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';

export const Register = () => {

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');



    const navigate = useNavigate();
    const hanldeRegisterSuccessfulClick = () => {
        navigate('/');
    }

    return (
        <div>
            <h1>Register</h1>
            <form>
                <label>Username:</label><br />
                <input type="text" /><br />
                <label>Password:</label><br />
                <input type="password" /><br />
                <button onClick={(e) => {
                    e.preventDefault();
                    hanldeRegisterSuccessfulClick();
                }}>Register</button>
            </form>
        </div >
    );
};