package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/axadjonovsardorbek/tender/config"
)

func SendMessage(message string) error {
	cf := config.LoadConfig()
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cf.BotToken)
	params := url.Values{}
	params.Add("chat_id", cf.GroupId)
	params.Add("text", message)

	// So'rov yuborish
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Javobni o'qish
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Status kodni tekshirish
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
