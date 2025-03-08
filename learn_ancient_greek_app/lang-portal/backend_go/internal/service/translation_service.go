package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// TranslationRequest represents a request to translate text
type TranslationRequest struct {
	Text string `json:"text"`
}

// TranslationResponse represents a response from the translation service
type TranslationResponse struct {
	Original   string `json:"original"`
	Translated string `json:"translated"`
	Error      string `json:"error,omitempty"`
}

// TranslationService handles communication with the translation model API
type TranslationService struct {
	ModelAPIURL string
}

// NewTranslationService creates a new instance of TranslationService
func NewTranslationService() *TranslationService {
	// Replace with your actual Lightning.ai API URL
	return &TranslationService{
		ModelAPIURL: "https://your-lightning-app-url.lightning.ai/translate",
	}
}

// Translate sends a translation request to the model API
func (s *TranslationService) Translate(text string) (TranslationResponse, error) {
	req := TranslationRequest{Text: text}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return TranslationResponse{Error: "Failed to process request"}, err
	}

	// Forward the request to the model API
	resp, err := http.Post(s.ModelAPIURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return TranslationResponse{Error: "Failed to connect to translation service"}, err
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return TranslationResponse{Error: "Failed to read response from translation service"}, err
	}

	// Parse the response
	var translationResponse TranslationResponse
	if err := json.Unmarshal(body, &translationResponse); err != nil {
		return TranslationResponse{Error: "Failed to parse response from translation service"}, err
	}

	return translationResponse, nil
}