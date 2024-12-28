import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import { jwtDecode } from "jwt-decode";
import {
  dropboxAuth,
  listUserFolders,
} from "../../services/dropboxService/dropboxService";

export const Home = () => {
  const [folders, setFolders] = useState([]);
  const [currentPath, setCurrentPath] = useState("");
  const [error, setError] = useState("");
  const [accessToken, setAccessToken] = useState("");

  useEffect(() => {
    const accessTokenCookie = cookies.find((cookie) =>
      cookie.startsWith("access_token=")
    );

    if (accessTokenCookie) {
      const token = accessTokenCookie.split("=")[1];
      const decodedToken = jwtDecode(token);
      setAccessToken(decodedToken.username);
    } else {
      console.log("Access token not found");
    }
  }, []);

  const cookieString = document.cookie;
  const cookies = cookieString.split("; ");

  const handleListFolders = async (path = "") => {
    try {
      const newPath = path.startsWith("/") ? path : "";
      const data = await listUserFolders(path, accessToken);

      console.log("response: ", data);
      console.log("p ", path);
      setFolders(data.folders.entries || []);
      setCurrentPath(newPath);
    } catch (err) {
      setError(err.response?.data?.error || "Failed to fetch folders");
    }
  };

  const handleDropboxClick = async () => {
    try {
      const authURL = await dropboxAuth();
      window.location.href = authURL;
    } catch (error) {
      console.error(error.message);
    }
  };

  const handleFolderClick = (folderName) => {
    const newPath = currentPath ? `${currentPath}/${folderName}` : folderName;
    handleListFolders(newPath);
  };

  const removeCookie = (cookie) => {
    document.cookie = `${cookie}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
  };

  //LOGOUT
  const handleLogOutUser = () => {
    try {
      const userToken = localStorage.getItem("authToken");

      if (!userToken) {
        console.log("User token not found");
        return;
      }
      localStorage.removeItem("authToken");
      removeCookie("access_token");
      window.location.reload();
    } catch (error) {
      console.error("Error during logout ", error.message);
    }
  };

  return (
    <>
      <Link to="/test1">Test </Link>
      <br />
      {!accessToken && <button onClick={handleDropboxClick}>Dropbox</button>}
      {accessToken && (
        <button onClick={() => handleListFolders("")}>
          List folders of user
        </button>
      )}
      <br />
      <h2>User folders</h2>
      <ul>
        {folders.map((folder, index) => (
          <li
            key={index}
            style={{ cursor: "pointer", color: "blue" }}
            onClick={() => handleFolderClick(folder.path_lower)}
          >
            {folder.name}
          </li>
        ))}
      </ul>
      <button onClick={handleLogOutUser}>Logout</button>
    </>
  );
};
