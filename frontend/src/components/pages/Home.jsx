import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { jwtDecode } from "jwt-decode";

export const Home = () => {
  const [folders, setFolders] = useState([]);
  const [error, setError] = useState("");
  const [token, setToken] = useState("");

  useEffect(() => {
    const accessTokenCookie = cookies.find((cookie) =>
      cookie.startsWith("access_token=")
    );

    if (accessTokenCookie) {
      const accessToken = accessTokenCookie.split("=")[1];
      console.log(accessToken);
      const decodedToken = jwtDecode(accessToken);
      console.log(decodedToken);
      console.log(decodedToken.username);
      setToken(decodedToken.username);
    } else {
      console.log("Access token not found");
    }
    const fetchFolders = async () => {
      try {
        const response = await axios.get("/api//folders", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setFolders(response.data.folders);
      } catch (err) {
        setError(err.response?.data?.error || "Failed to fetch folders");
      }
    };
    fetchFolders();
  }, []);

  const cookieString = document.cookie;

  const cookies = cookieString.split("; ");

  const handleDropboxClick = async () => {
    try {
      const response = await axios.get(
        "http://localhost:8080/api/dropbox_login"
      );

      if (response.status === 200 && response.data.authURL) {
        window.location.href = response.data.authURL;
      } else if (response.status !== 200) {
        console.error(`Unexpected status code: ${response.status}`);
        alert("Unexpected response from server. Please try again.");
      } else {
        console.error("Authorization URL not found in response.");
      }
    } catch (err) {
      console.error("Error initiating Dropbox login:", err.message);
      alert("Failed to initiate Dropbox login. Please try again.");
    }
  };

  return (
    <>
      <Link to="/test1">Test </Link>
      <br />
      <button onClick={handleDropboxClick}>Dropbox</button>
      <br />
      <h2>User folders</h2>
      <ul>
        {folders.map((folder, index) => {
          <li key={index}>{folder}</li>;
        })}
      </ul>
    </>
  );
};
