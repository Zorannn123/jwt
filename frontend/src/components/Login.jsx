import React from 'react'

export const Login = () => {
    return (
        <div>
            <h1>Login</h1>
            <form>
                <label>Username:</label><br />
                <input type="text" /><br />
                <label>Password:</label><br />
                <input type="password" /><br />
                <button>Login</button>
            </form>
        </div>
    );
};