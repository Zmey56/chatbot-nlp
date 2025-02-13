package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// Constants with URL OpenAI API
const OpenAIEndpoint = "https://api.openai.com/v1/completions"

type openAIClient struct {
	apiKey string
}

// NewOpenAIClient creates a new client for working with the OpenAI API
func NewOpenAIClient(apiKey string) NLPClient {
	return &openAIClient{apiKey: apiKey}
}

// SendRequest sends a text request to the OpenAI API and return a response
func (o *openAIClient) SendRequest(text string) (string, error) {
	requestBody, _ := json.Marshal(map[string]interface{}{
		"model":      "text-davinci-003",
		"prompt":     text,
		"max_tokens": 100,
	})

	req, _ := http.NewRequest("POST", OpenAIEndpoint, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+o.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error closing response body:", err)
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error from OpenAI API")
	}

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return "", err
	}

	choices, ok := res["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", errors.New("invalid response format")
	}

	firstChoice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", errors.New("unexpected response structure")
	}

	responseText, ok := firstChoice["text"].(string)
	if !ok {
		return "", errors.New("response text missing")
	}

	return responseText, nil
}
