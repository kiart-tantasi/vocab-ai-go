package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	apiKeyChatgpt string
}

type AIProvider int

const (
	CHATGPT AIProvider = iota
	GEMINI
)

func main() {
	loadEnvFiles()
	configs := newConfigs()
	testVocabs := []string{"irritate", "frustrate"}

	for _, e := range testVocabs {
		result, err := askAi(e, CHATGPT, configs.apiKeyChatgpt)
		if err != nil {
			fmt.Printf("Error when asking ChatGPT: %s\n", err)
		}
		fmt.Printf("[%s]\n", e)
		fmt.Println(result)
		fmt.Printf("=====================================================\n")
	}
}

func askAi(vocab string, provider AIProvider, apiKey string) (string, error) {
	switch provider {
	case CHATGPT:
		httpClient := newHttpClient()
		result, err := askChatGPT(vocab, httpClient, apiKey)
		if err != nil {
			return "", err
		}
		return result, nil
	case GEMINI:
		return "", fmt.Errorf("TODO: implement")
	default:
		return "", fmt.Errorf("TODO: implement")
	}
}

func newConfigs() *Config {
	return &Config{
		apiKeyChatgpt: os.Getenv("API_KEY_CHATGPT"),
	}
}

func askChatGPT(vocab string, httpClient *http.Client, apiKey string) (string, error) {
	url := "https://api.openai.com/v1/responses"
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{
		"model": "gpt-4o-mini",
		"input": "Can you give vocab defition and 2 - 3 daily usage examples for the word \"%s\""
	}`, vocab))
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func loadEnvFiles() {
	// Try .env.local and .env first
	err := godotenv.Load(".env.local", ".env")
	if err == nil {
		fmt.Println("Loaded .env and .env.local")
		return
	}

	// Then try .env only
	err = godotenv.Load(".env")
	if err == nil {
		fmt.Println("Loaded .env")
	}

	fmt.Printf("Error while loading .env: %s", err)
	os.Exit(1)
}
