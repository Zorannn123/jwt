import React from 'react'

export const Register = () => {
    return (
        <div>
            <h1>Register</h1>
            <form>
                <label>Username:</label><br />
                <input type="text" /><br />
                <label>Password:</label><br />
                <input type="password" /><br />
                <button>Register</button>
            </form>
        </div>
    );
};