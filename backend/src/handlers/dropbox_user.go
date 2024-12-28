package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCurrentUser(accessToken string) error {
	url := "https://api.dropboxapi.com/2/users/get_current_account"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer " + accessToken)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil{
		return fmt.Errorf("request failed: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK{
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("error: %v, body %s", response.StatusCode, body)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil{
		return fmt.Errorf("failed to decode response: %v", err)
	}

	//fmt.Printf("curr user %v\n", result)
	return nil
}

