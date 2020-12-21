package blizzardApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BlizzardApi struct {
	clientId     string
	clientSecret string
	Token        string `json:"access_token"`
	client       *http.Client
}

func Init(clientId string, clientSecret string) *BlizzardApi{
	return &BlizzardApi{
		clientId: clientId,
		clientSecret: clientSecret,
		client: &http.Client{},
	}
}

func (api *BlizzardApi) GetToken() {
	body := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest(http.MethodPost, "https://us.battle.net/oauth/token", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.SetBasicAuth(api.clientId, api.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := api.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	resData, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(resData, &api)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (api *BlizzardApi) setHeader (req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.Token))
}

func (api *BlizzardApi) MakeRequest (url string) []byte {
	api.GetToken()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	api.setHeader(req)

	res, err := api.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	resData, _ := ioutil.ReadAll(res.Body)
	return resData
}