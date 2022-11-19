package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

// NewBearerToken retrieves a valid bearerTokenProvider to be used by the client for authenticated requests.
func NewBearerToken(clientID, clientSecret string) (*securityprovider.SecurityProviderBearerToken, error) {
	var oauth2Config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		return nil, err
	}

	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken(token.AccessToken)
	if err != nil {
		return nil, err
	}
	return bearerTokenProvider, nil
}

type ValidateTokenResponse struct {
	ClientID  string   `json:"client_id"`
	Login     string   `json:"login"`
	Scopes    []string `json:"scopes"`
	UserID    string   `json:"user_id"`
	ExpiresIn int64    `json:"expires_in"`
}

// ValidateToken verifies that the access token is still valid.
func ValidateToken(token string) (*ValidateTokenResponse, error) {
	req, err := http.NewRequest("GET", "https://id.twitch.tv/oauth2/validate", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *ValidateTokenResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, err
}
