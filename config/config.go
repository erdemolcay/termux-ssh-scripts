package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
	"termux-ssh-scripts/util"
)

type Config struct {
	ApiToken string `json:"apiToken"`
	ZoneId   string `json:"zoneId"`
}

func New(subcommand string) *Config {
	c := Config{}

	switch subcommand {
	case "update":
		c.ApiToken, c.ZoneId = flags("update")
		if c.ApiToken == "" || c.ZoneId == "" {
			c = parse()
		}
	case "install":
		c.ApiToken, c.ZoneId = flags("install")
		if c.ApiToken == "" {
			log.Fatal("API token is required")
		}
		if c.ApiToken == "" {
			log.Fatal("Zone ID is required")
		}
	}
	return &c
}

func flags(command string) (apiToken string, zoneId string) {
	fs := flag.NewFlagSet(command, flag.ContinueOnError)
	a := fs.String("api-token", "", "Cloudflare API token")
	z := fs.String("zone-id", "", "Cloudflare zone ID")
	_ = fs.Parse(os.Args[2:])
	return *a, *z
}

func parse() (c Config) {
	f := util.HomeDir() + "/.config/termux-ssh-scripts/config.json"

	viper.SetConfigFile(f)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v\n", err)
	}

	return
}
