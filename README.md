# Env - env to struct
---
A simple and zero-dependencies library to parse environment variables into structs

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/env)](https://pkg.go.dev/github.com/go-zoox/env)
[![Build Status](https://github.com/go-zoox/env/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/env/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/env)](https://goreportcard.com/report/github.com/go-zoox/env)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/env/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/env?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/env.svg)](https://github.com/go-zoox/env/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/env.svg?label=Release)](https://github.com/go-zoox/env/releases)

Load application environment variables from a `.env` file into the current process.

```go
package main

import (
	"fmt"
	"time"

	"github.com/go-zoox/env"
)

type config struct {
	Home         string         `env:"HOME"`
	Port         int            `env:"PORT,default=8080"`
	Password     string         `env:"PASSWORD,unset"`
	IsProduction bool           `env:"PRODUCTION"`
	Hosts        []string       `env:"HOSTS,separator=;"`
	TempFolder   string         `env:"TEMP_FOLDER,default=/tmp"`
  Duration     time.Duration  `env:"DURATION"`
	StringInts   map[string]int `env:"MAP_STRING_INT"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

  // or you can use generics
  cfg, err := env.ParseAs[config]()
  if err != nil {
		fmt.Printf("%+v\n", err)
  }

	fmt.Printf("%+v\n", cfg)
}
```

## Inspired by
* [caarlos0/env](https://github.com/caarlos0/env)
* [joho/godotenv](https://github.com/joho/godotenv) - dot env file
* [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) - struct tag
* [JeremyLoy/config](https://github.com/JeremyLoy/config) - ci
