import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Login } from './components/auth/Login/Login';
import { Register } from './components/auth/Register/Register'
import { Home } from './components/pages/Home';
import { Test } from './components/pages/protecttest'


function App() {
  const [isLogged, setIsLogged] = useState(false);
  console.log(isLogged);
  const logState = () => {
    if (!isLogged) {
      setIsLogged(true);

    }
  };

  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login logState={logState} />} />
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<Register />} />

        {isLogged && <Route path="/test1" element={<Test />} />}




      </Routes>
    </Router>
  );
}

export default App;
