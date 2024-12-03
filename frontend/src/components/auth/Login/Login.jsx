import React, { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom';
import styles from './Login.module.css';

export const Login = ({ logState }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleLogin = async () => {
        try {
            const response = await axios.post('http://localhost:8080/api/login', {
                username: username,
                password: password
            });

            localStorage.clear();
            const token = response.data.token;
            if (token) {
                localStorage.setItem('userToken', token);

                setErrorMessage('');
                handleLoginSuccessfulClick();
                logState();
                window.alert('Login successful!');

            } else {
                console.error("Token not found in the response");
                setErrorMessage("Token not found in the response");
            }
        } catch (err) {
            console.error('Login failed:', errorMessage.response ? errorMessage.response.data : errorMessage.message);
            setErrorMessage('Login failed. Please check your credentials.');
        }
    };

    const navigate = useNavigate();
    const handleLoginSuccessfulClick = () => {
        navigate('/');
    }

    return (
        <div className={styles.back}>
            <div className={styles.container}>
                <form onSubmit={handleLogin}>
                    <h1 className={styles.test}>Login</h1>
                    <label>Username:</label><br />
                    <input type="text" value={username} onChange={(e) => setUsername(e.target.value)} /><br />
                    <label>Password:</label><br />
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} /><br />
                    {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
                    <button className={styles.btn} onClick={(e) => {
                        e.preventDefault();
                        if (username.trim() === '' || password.trim() === '') {
                            setErrorMessage('Please enter both username and password.');
                        } else {
                            handleLogin();
                            setErrorMessage('');
                        }
                    }}>Login</button>
                </form>
                <div>
                    Don't have an account?
                    <a href='/register' className={styles.link}> Sign up</a>
                </div>
            </div>
        </div>
    );
};