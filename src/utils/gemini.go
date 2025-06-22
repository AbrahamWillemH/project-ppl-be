package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func EvaluateEssayWithGemini(question, studentAnswer string, maxScore float64) (float64, error) {
	apiURL := os.Getenv("GEMINI_API_URL") + "?key=" + os.Getenv("GEMINI_API_KEY")

	prompt := fmt.Sprintf(`Evaluate the following essay answer:
Question: "%s"
Student Answer: "%s"

Assign a score from 0 to %.0f based on how well the answer satisfies the question.
Give a full score if highly relevant, partial if somewhat relevant, low if irrelevant.
Return only the score as a plain number.`,
		question, studentAnswer, maxScore)

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}
	bodyBytes, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return 0, fmt.Errorf("failed to create Gemini request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to call Gemini API: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to parse Gemini response: %w", err)
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return 0, fmt.Errorf("no candidates returned from Gemini")
	}

	var score float64
	_, err = fmt.Sscan(result.Candidates[0].Content.Parts[0].Text, &score)
	if err != nil {
		return 0, fmt.Errorf("failed to parse score: %w", err)
	}

	// Ensure score is within range
	if score > maxScore {
		score = maxScore
	}
	if score < 0 {
		score = 0
	}
	return score, nil
}
