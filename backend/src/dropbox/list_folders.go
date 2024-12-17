package dropbox

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListFolderRequest struct {
	Path string `json:"path"`
}

type ListFolderResponse struct {
	Entries []struct {
		Name string `json:"name"`
		Tag  string `json:".tag"`
	} `json:"entries"`
}

func ListUserFolders(accessToken string) ([]string, error) {
	client := &http.Client{}
	url := "https://api.dropboxapi.com/2/files/list_folder"

	requestBody := ListFolderRequest{Path: ""}
	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil{
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer " + accessToken)
	req.Header.Set("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	var listFolderResponse ListFolderResponse
	err = json.Unmarshal(body, &listFolderResponse)
	if err != nil{
		return nil, err
	}

	var folders []string
	for _, entry := range listFolderResponse.Entries{
		if entry.Tag == "folder"{
			folders = append(folders, entry.Name)
		}
	}

	return folders, nil 
}

func GetDropboxFolders(c *gin.Context){
	accessToken := c.GetHeader("Authorization")
	if accessToken == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "unathorized"})
		return
	}
	folders, err := ListUserFolders(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"folders" : folders})
}
