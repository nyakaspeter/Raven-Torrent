package v0

import "encoding/json"

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func successMessage() string {
	message := MessageResponse{
		Success: true,
		Message: "OK",
	}

	messageString, _ := json.Marshal(message)

	return string(messageString)
}
