package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Twitch_cred struct {
	gorm.Model
	ID uint 	`gorm:"primaryKey"`
	Token string `json:"access_token"`
	Expires_In string `json:"expires_in"`
}


func getToken() Twitch_cred {

	tClient := os.Getenv("TWITCH_CLIENT")
	tSecret := os.Getenv("TWITCH_SECRET")

	// url := "https://id.twitch.tv/oauth2/token?client_id=" + string(tClient) + "&client_secret=" + string(tSecret) + "&grant_type=client_credentials"

	jsonData := map[string]string{
		"client_id": tClient,
		"client_secret": tSecret,
		"grant_type": "client_credentials",
	}

	jsonValue, _ := json.Marshal(jsonData)

	request, _ := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed misareably %s \n", err)
	} 
	data, _ := ioutil.ReadAll(response.Body)
	// println(string(data))
	
	var tokenStruct Twitch_cred

	json.Unmarshal(data, &tokenStruct)

	println(tokenStruct.Token)

	return tokenStruct

}