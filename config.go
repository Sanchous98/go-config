package config

import (
	"errors"
	"github.com/Sanchous98/go-di"
	"gopkg.in/yaml.v3"
	"math"
	"os"
	"reflect"
)

const configTag = "config"
const configPath = "/config"

func init() {
	di.Application().Set(Configurator(&Config{}))
	di.Application().PostCompile(func(event di.Event) {
		container := event.GetElement().(di.Container)
		configurator := event.GetElement().(di.Container).Get(new(Configurator)).(Configurator)

		for _, service := range container.All() {
			_ = Unmarshall(configurator, service)
		}
	}, math.MinInt)
}

type Config struct {
	DotNotationBag
}

func (c *Config) Configure(service reflect.Value) error {
	if service.Kind() == reflect.Ptr {
		service = service.Elem()
	}

	for i := 0; i < service.NumField(); i++ {
		key, ok := service.Type().Field(i).Tag.Lookup(configTag)

		if !ok {
			continue
		}

		if value, err := c.Get(key); !errors.Is(err, &NotFoundError{}) {
			service.Field(i).Set(reflect.ValueOf(value))
		} else {
			return err
		}
	}

	return nil
}

func (c *Config) Constructor() {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + configPath)

	if err != nil {
		panic(err)
	}

	defer func() { _ = file.Close() }()

	c.Lock()
	_ = yaml.NewDecoder(file).Decode(c.DotNotationBag.bag)
	c.Unlock()
}

func Unmarshall(in Configurator, out interface{}) error {
	return in.Configure(reflect.ValueOf(out))
}
