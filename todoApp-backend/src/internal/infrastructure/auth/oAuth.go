package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type GoogleTokenResponse struct {
	AccessToken string `json:"access_token"`
	IdToken     string `json:"id_token"`
}

type GoogleUser struct {
	Email string `json:"email"`
	Name  string `json:"username"`
}

func GetAccessToken(code string) (*GoogleTokenResponse, error) {
	if code == "" {
		return nil, errors.New("no code provided")
	}

	resp, err := http.PostForm("\"https://oauth2.googleapis.com/token", map[string][]string{
		"client_id":     {os.Getenv("GOOGLE_CLIENT_ID")},
		"client_secret": {os.Getenv("GOOGLE_CLIENT_SECRET")},
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {os.Getenv("GOOGLE_REDIRECT_URI")},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var tokenResponse GoogleTokenResponse

	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

func GetUserInfo(AccessToken string) (string, string, error) {
	if AccessToken == "" {
		return "", "", errors.New("no access token provided")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + AccessToken)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	var userInfo GoogleUser

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return "", "", err
	}

	return userInfo.Name, userInfo.Email, nil
}
