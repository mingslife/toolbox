package conf

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var cfg *Config
var wg sync.WaitGroup

type Config struct {
	Node  int64 `default:"0"`
	Debug bool  `default:"true"`
	Http  Http
	Grpc  Grpc
}

func ParseConfig() *Config {
	wg.Add(1)
	defer wg.Done()

	cfg = &Config{}
	if err := envconfig.Process("toolbox", cfg); err != nil {
		log.Fatalln(err)
	}
	return cfg
}

func GetConfig() *Config {
	wg.Wait()
	return cfg
}
