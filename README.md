# Go Configurator

Easily configure services, using [Go DI Container](https://github.com/Sanchous98/go-di)

```yaml
another:
    test: value
```

```go
package main

type TestStruct struct {
    TestField string `config:"another.test"` // value
}
```

Config is loaded from "%WORKING_DIRECTORY%/config"

To replace the default configurator, just replace the configurator in container

```go
func init() {
    di.Application().Set(config.Configurator(&YourConfigurator{}))
}
```
