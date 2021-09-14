package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type InitBody struct {
	Type string `json:"type"`
	Version string `json:"version"`
	Condition Condition `json:"condition"`
	Transport Transport `json:"transport"`
}

type Condition struct {
	BroadcasterUserID string `json:broadcaster_user_id`
}

type Transport struct {
	Method string `json:"method"`
	Callback string `json:"callback"`
	Secret string `json:"secret"`
}

func initEventSub() {
	Rclient := resty.New()
	tClient := os.Getenv("TWITCH_CLIENT")

	condition := Condition{
		BroadcasterUserID: "369724",
	}

	transport := Transport {
		Method: "webhook",
		Callback: "https://8fa8-99-33-185-199.ngrok.io/callback",
		Secret: "testwoo",
	}

	jsonData := InitBody {
		Type: "channel.update",
		Version: "1",
		Condition: condition,
		Transport: transport,
	}
	
	jsonValue, _ := json.Marshal(jsonData)

	token := "Bearer " + getToken().Token

	resp, err := Rclient.R().
      SetHeader("Content-Type", "application/json").
	  SetHeader("Authorization", token).
	  SetHeader("Client-ID", tClient).
      SetBody(jsonValue).
	  Post("https://api.twitch.tv/helix/eventsub/subscriptions")
                                                              
	if err != nil {
		fmt.Printf("The init request failed misareably %s \n", err)
	} 

		
	println(resp)

}

func handleCallback(c *fiber.Ctx) {

}