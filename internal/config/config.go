package config

import (
	"flag"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

const (
	DefaultHost            = "localhost:8080"
	DefaultKey             = "13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b"
	DefaultDataBaseAddress = "host=localhost port=5432 user=postgres password=12345 dbname=hw-5 sslmode=disable"
)

// Config конфигурация
type Config struct {
	Host            string `env:"RUN_ADDRESS" envDefault:"localhost:8080"`
	Key             string `env:"DATABASE_URI" envDefault:"13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b"`
	DataBaseAddress string `env:"DATABASE_URI" envDefault:"host=localhost port=5432 user=postgres password=12345 dbname=hw-5 sslmode=disable"`
	TokenDuration   time.Duration
}

func (c *Config) parseArgsCMD() {
	if !flag.Parsed() {
		flag.StringVar(&c.Host, "a",
			DefaultHost, "grpc server launching address")
		flag.StringVar(&c.Key, "k", DefaultKey,
			"security key")
		flag.StringVar(&c.DataBaseAddress, "d", DefaultDataBaseAddress,
			"Address of db connection")
		flag.Parse()
	}
}

func (c *Config) applyConfig(other Config) {
	if c.Host == DefaultHost {
		c.Host = other.Host
	}
	if c.Key == DefaultKey {
		c.Key = other.Key
	}
	if c.DataBaseAddress == DefaultDataBaseAddress {
		c.DataBaseAddress = other.DataBaseAddress
	}
}

func (c *Config) Init() error {
	log.Println("Init config")
	var c2 Config
	//parsing env config
	err := env.Parse(c)
	if err != nil {
		return err
	}
	//parsing command line config
	c2.parseArgsCMD()
	//applying config
	c.applyConfig(c2)
	return nil
}
