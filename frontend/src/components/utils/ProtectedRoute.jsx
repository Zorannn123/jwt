import React from "react";
import { Navigate, Outlet } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';

const isTokenValid = () => {
    const token = localStorage.getItem('userToken');
    if (!token) {
        return console.error("Wrong token");
    };

    try {
        const decodedToken = jwtDecode(token);
        const currTime = new Date().getTime() / 1000;
        if (decodedToken.exp < currTime) {
            localStorage.clear();
            return false;
        }
    } catch (error) {
        console.error("Failed to decode token", error);
        return false;
    }
    return true;
};

export const ProtectedRoute = () => {
    const isLoggedIn = isTokenValid();
    return (
        isLoggedIn ? <Outlet /> : <Navigate to='/' />
    )
};
