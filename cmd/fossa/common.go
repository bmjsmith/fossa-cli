package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Common utilities among commands
func makeAPIRequest(requestType string, url string, payload []byte, apiKey string) ([]byte, error) {
	req, err := http.NewRequest(requestType, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("invalid API key (check the $FOSSA_API_KEY environment variable)")
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad server response (" + string(resp.StatusCode) + ")")
	}
	return ioutil.ReadAll(resp.Body)
}
