package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type TokenRes struct {
	Token   string `json:"access_token"`
	Scope   string `json:"scope,omitempty"`
	Expires int64  `json:"expires_in"`
	Type    string `json:"token_type"`
}

type Req struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
}

type Client struct {
	TokenRes
	TimeRequested time.Time `json:"time_requested"`
}

var clients map[string]Client = make(map[string]Client)

// GetToken is the service function to get a JWT from Auth0
// using valid client id, client secret and audience fields.
func (s service) GetToken(decoder *json.Decoder) (*Client, error) {
	var currReq Req
	err := decoder.Decode(&currReq)
	if err != nil {
		return nil, err
	}

	if client, present := clients[currReq.ClientId]; present {
		newExpires := time.Now().UnixMilli()/1000 - client.TimeRequested.UnixMilli()/1000
		if newExpires < int64(client.Expires)-10 {
			client.Expires = int64(client.Expires - newExpires)
			return &client, nil
		}
	}

	url := "https://dev--180yk9n.us.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"" + currReq.ClientId + "\",\"client_secret\":\"" + currReq.ClientSecret + "\",\"audience\":\"" + currReq.Audience + "\",\"grant_type\":\"client_credentials\"}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r TokenRes
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	newClient := Client{r, time.Now()}
	clients[currReq.ClientId] = newClient

	return &newClient, nil
}
