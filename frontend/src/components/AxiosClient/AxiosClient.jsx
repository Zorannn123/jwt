import axios from "axios";

const API_BASE_URL = process.env.REACT_APP_API_BASE_URL;

const apiClient = axios.create({
  baseURL: `${API_BASE_URL}`,
});

apiClient.interceptors.request.use(
  (config) => {
    if (!config.headers["Authorization"]) {
      const authToken = localStorage.getItem("authToken");
      if (authToken) {
        config.headers["Authorization"] = `Bearer ${authToken}`;
      }
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

apiClient.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response && error.response.status === 401) {
      localStorage.removeItem("authToken");

      window.location.reload();
    }

    return Promise.reject(error);
  }
);

export default apiClient;
