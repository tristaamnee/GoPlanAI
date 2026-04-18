package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goplan-ai/internal/models"
	"log"
	"net/http"
	"time"
)

type GroqAI struct {
	APIKey string
}

func (g *GroqAI) GenerateTasks(title, description string) ([]models.Task, error) {
	prompt := fmt.Sprintf(
		"Short Task List for: %s. Description: %s. Return ONLY a plain JSON array of 3 objects. No talk, no markdown.",
		title, description,
	)
	url := "https://api.groq.com/openai/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "llama3-70b-8192",
		"messages": []map[string]interface{}{
			{"role": "user", "content": prompt},
		},
		"response_format": map[string]string{"type": "json_object"},
	}

	jsonData, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+g.APIKey)

	client := &http.Client{
		Timeout: time.Second * 90,
	}

	log.Println("Đang gọi Groq API...")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lỗi kết nối Groq: %v", err)
	}
	defer resp.Body.Close()
	log.Println("Groq đã phản hồi!")

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Groq API sending error: %s", resp.Status)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("lỗi decode JSON: %v", err)
	}
	if len(result.Choices) == 0 {
		return nil, fmt.Errorf("No choices returned")
	}

	content := result.Choices[0].Message.Content
	var tasks []models.Task
	err = json.Unmarshal([]byte(content), &tasks)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling tasks: %s", err)
	}

	return tasks, err
}
