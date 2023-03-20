package main

import "github.com/afex/hystrix-go/hystrix"

func SetupCb() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               5000,
		MaxConcurrentRequests: 1000,
		ErrorPercentThreshold: 20,
	})
}
