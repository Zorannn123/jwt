package dropbox

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		PathDisplay string `json:"path_display"`
		PathLower string `json:"path_lower"`
	} `json:"entries"`
	Cursor string `json:"cursor"`
	HasMore bool `json:"has_more"`
}

func ListUserFolders(accessToken string, path string) (*ListFolderResponse, error) {
	client := &http.Client{}
	url := "https://api.dropboxapi.com/2/files/list_folder"

	requestBody := ListFolderRequest{Path: path}
	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil{
		return nil, err
	}

	req.Header.Set("Authorization",  accessToken)
	req.Header.Set("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	fmt.Println(body)

	var listFolderResponse ListFolderResponse
	err = json.Unmarshal(body, &listFolderResponse)
	if err != nil{
		return nil, err
	}
	fmt.Println(listFolderResponse)

	// var folders []string
	// for _, entry := range listFolderResponse.Entries{
	// 	fmt.Println("path ", entry.PathDisplay)
	// 	if entry.Tag == "folder"{
	// 		fmt.Println(entry.PathDisplay)
	// 		folders = append(folders, entry.Name)
	// 	}
	// }

	for listFolderResponse.HasMore{
		cursorRequestBody := map[string]string{"cursor": listFolderResponse.Cursor}
		cursorJsonBody, _ := json.Marshal(cursorRequestBody)
		fmt.Println("sss ", cursorJsonBody)

		cursorReq, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/files/list_folder/continue", bytes.NewBuffer(cursorJsonBody))
		if err != nil {
			return nil, err
		}

		cursorReq.Header.Set("Authorization", accessToken)
		cursorReq.Header.Set("Content-type", "application/json")
		cursorResp, err := client.Do(cursorReq)
		if err != nil {
			return nil, err
		}
		defer cursorResp.Body.Close()

		cursorBody, err := ioutil.ReadAll(cursorResp.Body)
		if err != nil {
			return nil, err
		}

		var cursorResponse ListFolderResponse
		err = json.Unmarshal(cursorBody, &cursorResponse)
		if err != nil {
			return nil, err
		}

		// for _, entry := range cursorResponse.Entries {
		// 	if entry.Tag == "folder" {
		// 		folders = append(folders, entry.Name)
		// 	}
		// }

		// for _, entry := range cursorResponse.Entries {
		// 	if entry.Tag == "file" {
		// 		files = append(folders, entry.Name)
		// 	}
		// }
		listFolderResponse.HasMore = cursorResponse.HasMore
		listFolderResponse.Cursor = cursorResponse.Cursor
	}

	return &listFolderResponse, nil 
}

func GetDropboxFolders(c *gin.Context, path string){
	accessToken := c.GetHeader("Authorization")
	fmt.Println(accessToken)
	if accessToken == ""{
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "unathorized"})
		return
	}
	folders, err := ListUserFolders(accessToken, path)
	fmt.Println("fold \n", folders)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"folders" : folders, "path" : path})
}