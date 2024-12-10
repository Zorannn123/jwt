import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";

export const Home = () => {
  const [error, setError] = useState(null);

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const code = params.get("code");

    if (code) {
      axios
        .get("http://localhost:8080/api/auth/callback", { code })
        .then((response) => {
          console.log("Dropbox login successful", response.data);
          localStorage.setItem("dropbox_token", response.data.access_token);
        })
        .catch((err) => {
          console.error("Error during Dropbox login callback", err);
          setError("Failed to complete Dropbox login. Please try again.");
        });
    }
  }, []);

  return (
    <>
      <h1>Welcome to the Home Page</h1>
      {error && <p style={{ color: "red" }}>{error}</p>}
      <Link to="/test1">Test Page</Link>
      <br />
    </>
  );
};
