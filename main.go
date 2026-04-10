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
	"strings"

	"github.com/BurntSushi/toml"
)

// Constants
var (
	BaseUrl        = "https://api.telegram.org"
	Command        = "sendMessage"
	ParseMode      = "HTML"
	DefaultMessage = "Ding!"
)

// Config
type Config struct {
	Telegram struct {
		Bot_token string `toml:"token"`
		Chat_ID   string `toml:"chat_id"`
	} `toml:"telegram"`
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
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func sendMessage(config *Config, m SendMessageDTO) error {
	base, err := url.Parse(BaseUrl)
	if err != nil {
		return errors.New("internal error")
	}

	base = base.JoinPath("bot" + config.Telegram.Bot_token).JoinPath(Command)

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

func GetMessage() string {
	if len(os.Args) == 0 {
		return DefaultMessage
	}

	return strings.Join(os.Args[1:], " ")

}

// Entrypoint
func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	msg := SendMessageDTO{
		config.Telegram.Chat_ID,
		GetMessage(),
		ParseMode,
	}

	err = sendMessage(config, msg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
