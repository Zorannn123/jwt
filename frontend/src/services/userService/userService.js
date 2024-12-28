import apiClient from "../../components/AxiosClient/AxiosClient";

export const loginUser = async (username, password) => {
  try {
    const response = await apiClient.post("/api/login", {
      username,
      password,
    });
    return response.data;
  } catch (error) {
    console.error("Error during login: ", error.message);
    throw new Error("Failed to login.");
  }
};

export const registerUser = async (username, password) => {
  try {
    const response = await apiClient.post("/api/register", {
      username,
      password,
    });
    return response.data;
  } catch (error) {
    if (!error.response) {
      console.error("Network error:", error.message);
      throw new Error("Network error. Please check your connection.");
    }
    throw error;
  }
};
