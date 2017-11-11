package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/rs/zerolog/log"
	"github.com/sethgrid/pester"
)

// NikePlusAPIClient is the interface for the client that communicates with the Nike+ API
type NikePlusAPIClient interface {
	GetAccessToken(string, string, string) (NikeAccessToken, error)
}

type nikePlusAPIClientImpl struct {
}

// NewNikePlusAPIClient returns a new NikePlusAPIClient
func NewNikePlusAPIClient() (nikePlusAPIClient NikePlusAPIClient) {

	nikePlusAPIClient = &nikePlusAPIClientImpl{}

	return
}

// NikeAccessToken is the response returned by GetAccesToken
type NikeAccessToken struct {
	UserID      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// AccessTokenRequestBody is the request body sent by GetAccessToken
type AccessTokenRequestBody struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	KeepMeLoggedIn bool   `json:"keepMeLoggedIn"`
	ClientID       string `json:"client_id"`
	UXID           string `json:"ux_id"`
	GrantType      string `json:"grant_type"`
}

func (cl *nikePlusAPIClientImpl) GetAccessToken(email, password, clientID string) (nat NikeAccessToken, err error) {

	rb := AccessTokenRequestBody{
		Username:       email,
		Password:       password,
		KeepMeLoggedIn: true,
		ClientID:       clientID,
		UXID:           "com.nike.commerce.nikedotcom.web",
		GrantType:      "password",
	}

	// serialize json request body
	var requestBody io.Reader
	data, err := json.Marshal(rb)
	if err != nil {
		return
	}
	requestBody = bytes.NewReader(data)

	// generate uuid
	u4, err := uuid.NewV4()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%v%v", "https://unite.nike.com/loginWithSetCookie?appVersion=315&experienceVersion=276&uxid=com.nike.commerce.nikedotcom.web&locale=en_US&backendEnvironment=identity&browser=Google%20Inc.&os=undefined&mobile=false&native=false&visit=1&visitor=", u4)

	// call api
	response, err := pester.Post(url, "application/json", requestBody)
	if err != nil {
		return
	}

	defer response.Body.Close()

	// read response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	// unmarshal json body
	var b interface{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Error().Err(err).
			Interface("responseHeaders", response.Header).
			Str("responseBody", string(body)).
			Msg("Deserializing response from loginWithSetCookie failed")
		return
	}
	log.Debug().Interface("responseBody", b).Msg("Response for GetAccessToken")

	// unmarshal json body
	err = json.Unmarshal(body, &nat)
	if err != nil {
		return
	}

	return
}
