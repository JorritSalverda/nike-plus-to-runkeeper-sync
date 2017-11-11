package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverrideEnvvars(t *testing.T) {

	t.Run("GetAccessTokenReturnsValidAccessToken", func(t *testing.T) {

		nikePlusAPIClient := NewNikePlusAPIClient()
		testSecrets, err := readTestSecrets()

		// act
		nat, err := nikePlusAPIClient.GetAccessToken(testSecrets.Username, testSecrets.Password, testSecrets.ClientID)

		assert.NotNil(t, nat)
		assert.Nil(t, err)
		assert.Equal(t, testSecrets.UserID, nat.UserID)
	})
}

type TestSecrets struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID string `json:"client_id"`
	UserID   string `json:"user_id"`
}

func readTestSecrets() (testSecrets TestSecrets, err error) {

	data, err := ioutil.ReadFile("testSecrets.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &testSecrets)

	return
}
