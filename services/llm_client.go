package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var LLM_API_BASE = os.Getenv("LLM_URL") // or your deployed URL

type SetContextRequest struct {
	Context   map[string]interface{} `json:"context"`
	SessionID string                 `json:"session_id,omitempty"`
}
type SetContextResponse struct {
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
}
type ChatRequest struct {
	Message   string `json:"message"`
	SessionID string `json:"session_id"`
}
type ChatResponse struct {
	Response  string                   `json:"response"`
	SessionID string                   `json:"session_id"`
	History   []map[string]interface{} `json:"history"`
}

func SetContext(context map[string]interface{}, sessionID string) (SetContextResponse, error) {
	req := SetContextRequest{Context: context, SessionID: sessionID}
	body, _ := json.Marshal(req)
	resp, err := http.Post(LLM_API_BASE+"/set_context", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return SetContextResponse{}, err
	}
	defer resp.Body.Close()
	var result SetContextResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func Chat(message, sessionID string) (ChatResponse, error) {
	req := ChatRequest{Message: message, SessionID: sessionID}
	body, _ := json.Marshal(req)
	resp, err := http.Post(LLM_API_BASE+"/chat", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return ChatResponse{}, err
	}
	defer resp.Body.Close()
	var result ChatResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func GetHistory(sessionID string) ([]map[string]interface{}, error) {
	resp, err := http.Get(fmt.Sprintf("%s/history/%s", LLM_API_BASE, sessionID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result struct {
		SessionID string                   `json:"session_id"`
		History   []map[string]interface{} `json:"history"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.History, nil
}

func CheckHealth() (map[string]interface{}, error) {
	resp, err := http.Get(LLM_API_BASE + "/health")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result, nil
}
