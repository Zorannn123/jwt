import React, { useEffect, useState } from "react";
import {
  BrowserRouter as Router,
  Route,
  Routes,
  useNavigate,
} from "react-router-dom";
import { Login } from "./components/auth/Login/Login";
import { Register } from "./components/auth/Register/Register";
import { Home } from "./components/pages/Home";
import { Test } from "./components/pages/protecttest";
import { ProtectedRoute } from "./components/utils/ProtectedRoute";
import axios from "axios";

function App() {
  const [isLogged, setIsLogged] = useState(
    localStorage.getItem("authToken") != null
  );
  const token = localStorage.getItem("authToken");

  const logState = () => {
    setIsLogged(true);
  };
  //logout
  const logOut = () => {
    localStorage.clear();
    setIsLogged(false);
  };

  useEffect(() => {
    if (token) {
      axios
        .get("http://localhost:8080/api/secured/me", {
          headers: { Authorization: token },
        })
        .then((res) => setIsLogged(res.status === 200))
        .catch((err) => {
          console.error(err);
          localStorage.removeItem("authToken");
          setIsLogged(false);
        });
    } else {
      setIsLogged(false);
    }
  }, []);

  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login logState={logState} />} />
        <Route path="/register" element={<Register />} />

        <Route element={<ProtectedRoute isLoggedIn={isLogged} />}>
          <Route path="/" element={<Home />} />
          <Route path="/test1" element={<Test />} />
        </Route>
      </Routes>
    </Router>
  );
}

export default App;
