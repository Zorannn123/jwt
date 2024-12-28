import axios from "axios";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { registerUser } from "../../../services/userService/userService";

export const Register = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const handleRegister = async () => {
    try {
      const data = await registerUser(username, password);
      console.log(data);
      setErrorMessage("");
      hanldeRegisterSuccessfulClick();
    } catch (error) {
      if (error.response && error.response.status === 409) {
        setErrorMessage("Username already exists.");
        return;
      } else {
        setErrorMessage("Registration failed. Please check your credentials.");
      }
    }
  };

  const navigate = useNavigate();
  const hanldeRegisterSuccessfulClick = () => {
    navigate("/login");
  };

  return (
    <div>
      <h1>Register</h1>
      {errorMessage && <p style={{ color: "red" }}>{errorMessage}</p>}
      <form>
        <label>Username:</label>
        <br />
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <br />
        <label>Password:</label>
        <br />
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <br />
        <button
          onClick={(e) => {
            e.preventDefault();
            if (username.trim() === "" || password.trim() === "") {
              setErrorMessage("Please fill all fields!");
            } else {
              handleRegister();
              setErrorMessage("");
            }
          }}
        >
          Register
        </button>
      </form>
    </div>
  );
};
