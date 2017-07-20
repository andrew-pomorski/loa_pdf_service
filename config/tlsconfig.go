package config

import (
    "encoding/json"
    "log"
)


type TlsConfig struct {
    Servercert string
    Serverkey string
    Port string
}

func GetTlsConfig (dir string) TlsConfig {
    file := Open("tls.json", dir)
    defer Close(file);
    decoder := json.NewDecoder(file)
    config := TlsConfig{};
    err := decoder.Decode(&config)
    if err != nil {
        log.Println("error:", err)
        panic("Failed to read config tls")
    }
    return config
}
func (cfg TlsConfig)GetCert()string{
    return cfg.Servercert
}
func (cfg TlsConfig)GetKey()string{
    return cfg.Serverkey
}
func (cfg TlsConfig)GetPort()string{
    return cfg.Port
}

