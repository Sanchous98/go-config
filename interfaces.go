package config

import "reflect"

type ParameterBag interface {
	Get(string) (interface{}, error)
	Has(string) bool
	Set(string, string)
	Load(func(interface{}) error) error
}

type Configurator interface {
	Configure(reflect.Value) error
}
