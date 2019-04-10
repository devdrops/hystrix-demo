package main

import (
	"fmt"
	// Don't forget to `go get github.com/afex/hystrix-go/hystrix` ;)
	"github.com/afex/hystrix-go/hystrix"
	// Don't forget to `go get github.com/valyala/fasthttp` ;)
	"github.com/valyala/fasthttp"
)

func doRequest(url string) (error, string) {
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)

	response := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(request, response)
	if err != nil {
		return err, ""
	}

	return nil, string(response.Body())
}

func doTheThing() {
	hystrix.ConfigureCommand("hystrix_demo", hystrix.CommandConfig{
		Timeout:		1000,
		MaxConcurrentRequests:	100,
		ErrorPercentThreshold:	25,
	})

	hystrix.Go("hystrix_demo", func() error {
		for counter := 0; counter < 10; counter++ {
			if counter == 5 {
				err, response := doRequest("https://hystrix-demo.free.beeceptor.com/check?status=0")
				if err != nil {
					return err
				}
				fmt.Println(response)
			}
			err, response := doRequest("https://hystrix-demo.free.beeceptor.com/check?status=1")
			if err != nil {
				return err
			}
			fmt.Println(response)
		}

		return nil
	}, func(err error) error {
		err, response := doRequest("https://hystrix-demo.free.beeceptor.com/foo")
		fmt.Println(response)

		return nil
	})
}

func main() {
	doTheThing()
}
