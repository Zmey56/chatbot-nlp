package endpoint

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zmey56/chatbot-nlp/internal/service"
)

type ChatHandler struct {
	botService service.BotService
}

// NewChatHandler creates a new instance of chat handler
func NewChatHandler(botService service.BotService) *ChatHandler {
	return &ChatHandler{botService: botService}
}

func (h *ChatHandler) HandleChatRequest(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Message string `json:"message"`
	}

	// Decode the request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Error decoding request:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Println("Received message:", request.Message)

	// Get response from the bot service
	response, err := h.botService.GetResponse(request.Message)
	if err != nil {
		log.Println("Error processing message:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Sending response:", response)

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": response})
}
