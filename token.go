package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TokenKey struct {
	Alg   string `json:"alg"`
	Value string `json:"value"`
}

func GetTokenKey(tokenKeyURL string, clientID string, clientSecret string) (key *TokenKey, err error) {
	req, err := http.NewRequest(http.MethodGet, tokenKeyURL, nil)
	if err != nil {
		return
	}
	req.SetBasicAuth(clientID, clientSecret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, resp.Body.Close())
	}()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("invalid request/invalid response, code %d", resp.StatusCode)
		return
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var values map[string]any
	if err = json.Unmarshal(data, &values); err != nil {
		return
	}
	key = &TokenKey{
		Alg:   values["alg"].(string),
		Value: values["value"].(string),
	}
	return
}
