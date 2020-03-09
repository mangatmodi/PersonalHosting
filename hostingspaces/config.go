package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var (
	ErrInvalidPort  = errors.New("Invalid Port")
	ErrInvalidRoute = errors.New("Invalid Route/s")
	ErrNoRoutes     = errors.New("No Routes")
)

type Config struct {
	Port   uint16
	Routes []Route
}

type Route struct {
	FromSpaces string
	ToRoute    string
}
type yamlConfig struct {
	Port   uint16   `yaml:"port"`
	Routes []string `yaml:"routes"`
}

func (y *yamlConfig) toConfig() (*Config, error) {
	if (y.Port == 0) {
		return nil, ErrInvalidPort
	}
	r := make([]Route, 0, len(y.Routes))
	for _, el := range y.Routes {
		arr := strings.Split(el, " ")
		if len(arr) != 2 {
			return nil, ErrInvalidRoute
		}
		r = append(r, Route{
			FromSpaces: arr[0],
			ToRoute:    arr[1],
		})
	}
	if len(r) == 0 {
		return nil, ErrNoRoutes
	}

	return &Config{
		Port:   y.Port,
		Routes: r,
	}, nil
}

func NewConfig(path string) (*Config, error) {
	config := &yamlConfig{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config.toConfig()
}
