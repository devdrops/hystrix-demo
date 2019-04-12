package main

import (
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	hystrix.Go("basic_command", func() error {
		
		return nil
	}, nil)
}
