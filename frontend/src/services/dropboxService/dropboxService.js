import apiClient from "../../components/AxiosClient/AxiosClient";

export const dropboxAuth = async () => {
  try {
    const response = await apiClient.get("/api/dropbox_login");

    if (response.status === 200 && response.data.authURL) {
      return response.data.authURL;
    } else if (response.status !== 200) {
      throw new Error(`Unexpected status code: ${response.status}`);
    } else {
      throw new Error("Authorization URL not found in response.");
    }
  } catch (error) {
    console.error("Error during dropbox authorization: ", error.message);
    throw new Error("Failed dropbox authorization.");
  }
};

export const listUserFolders = async (path, accessToken) => {
  try {
    const response = await apiClient.get(`/api/folders?path=${path}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || "Failed to fetch folders";
  }
};
