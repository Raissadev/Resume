package utils

import (
	"encoding/json"
	"net/http"
)

type MessageRequest struct {
	//
}

type MessageRequestInterface interface {
	WriteMessage()
}

func (ms *MessageRequest) WriteMessage(w http.ResponseWriter, message string, status int) (int, error) {
	w.WriteHeader(status)

	now := map[string]interface{}{
		"message": message,
	}

	data, err := json.Marshal(now)

	if err != nil {
		return 0, err
	}

	return w.Write(data)
}

func (ms *MessageRequest) WriteBodyMessage(w http.ResponseWriter, message error, status int) (int, error) {
	w.WriteHeader(status)

	now := map[string]interface{}{
		"data": message,
	}

	data, err := json.Marshal(now)

	if err != nil {
		return 0, err
	}

	return w.Write(data)
}
