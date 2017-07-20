package config

import (
    "encoding/json"
    "log"
)


type MailConfig struct {
    Username string
    Password string
    Host string
    Port int
    Tolist []string
    From string
}

func GetMailConfig (dir string) MailConfig {
    file := Open("mail.json", dir)
    defer Close(file);
    decoder := json.NewDecoder(file)
    config := MailConfig{};
    err := decoder.Decode(&config)
    if err != nil {
        log.Println("error:", err)
        panic("Failed to read config mail")
    }
    return config
}

