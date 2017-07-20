package config

import (
	"os"
	"log"
)

type Config struct {
	path string
}

var cfg *Config= nil

func Setup(path string){
	cfg = &Config{path: path}
}

func Open(cfgname string, dir string) *os.File {
	if cfg == nil {
		file, err := os.Open(dir + "/" + cfgname)
		if err != nil {
			log.Fatal(err)
		}
		return file
	}else{
		dir = cfg.path
		file, err := os.Open(dir + "/" + cfgname)
		if err != nil {
			log.Fatal(err)
		}
		return file
	}
}

func Close(file *os.File){
	file.Close();
}
