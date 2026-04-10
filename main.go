package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Constants
var (
	BaseUrl        = "https://api.telegram.org"
	Command        = "sendMessage"
	DefaultMessage = "Ding!"
)

// Config
type Config struct {
	Bot_token string
	Chat_ID   string
}

func LoadConfig() (*Config, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	confPath := filepath.Join(confDir, "tg_notify", "conf.toml")

	conf, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("Couldn't load config at %s", confPath)
	}

	var config Config
	err = toml.Unmarshal(conf, &config)
	if err != nil {
		return nil, fmt.Errorf("Invalid conf.toml %s", &config)
	}

	return &config, nil
}

// Functions
type SendMessageDTO struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessage(config *Config, m SendMessageDTO) error {
	base, err := url.Parse(BaseUrl)
	if err != nil {
		return errors.New("internal error")
	}

	base = base.JoinPath("bot" + config.Bot_token).JoinPath(Command)

	jsonData, err := json.Marshal(m)
	if err != nil {
		return errors.New("internal error")
	}

	body := bytes.NewBuffer(jsonData)
	resp, err := http.Post(base.String(), "application/json", body)
	if err != nil {
		return errors.New("internal error")
	}
	defer resp.Body.Close()

	return nil

}

// Entrypoint
func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	msg := SendMessageDTO{
		config.Chat_ID,
		DefaultMessage,
	}

	err = sendMessage(config, msg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
