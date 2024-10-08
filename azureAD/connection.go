package azureAD

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vmwarepivotallabs/cf-mgmt/config"
	"github.com/xchapter7x/lo"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Token struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func getAccessToken(config *config.AzureADConfig) (Token, error) {
	token := Token{}
	reqBody := url.Values{}
	reqBody.Set("client_id", config.ClientId)
	reqBody.Set("client_secret", config.Secret)
	reqBody.Set("scope", Scope)
	reqBody.Set("grant_type", GrantType)

	tokenUrl := TokenLoginURL + config.TenantID + TokenOAuthURI
	req, err := http.NewRequest(http.MethodPost, tokenUrl, strings.NewReader(reqBody.Encode()))
	if err != nil {
		return token, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return token, err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return token, errors.New(fmt.Sprintf("response %d form call to %s: %s", response.StatusCode, tokenUrl, body))
	}

	// Result processing

	if err = json.Unmarshal(body, &token); err != nil {
		lo.G.Debugf("Reading body into Token struct failed: %s", err)
		return token, nil
	}
	return token, nil
}
