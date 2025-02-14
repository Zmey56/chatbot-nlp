package service

import (
	"errors"
	"log"

	"github.com/Zmey56/chatbot-nlp/internal/client"
)

type BotService interface {
	GetResponse(input string) (string, error)
}

type botService struct {
	apiClient client.NLPClient
}

// NewBotService creates a new instance BotService
func NewBotService(apiClient client.NLPClient) BotService {
	return &botService{apiClient: apiClient}
}

// GetResponse gets a response from the NLP API
func (b *botService) GetResponse(input string) (string, error) {
	if input == "" {
		return "", errors.New("empty input message")
	}

	response, err := b.apiClient.SendRequest(input)
	if err != nil {
		log.Println("Error getting response from NLP API:", err)
		return "", err
	}

	log.Println("Received response from NLP API:", response)

	return response, nil
}
