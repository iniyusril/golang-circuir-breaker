package main

import (
	"errors"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/sirupsen/logrus"
)

func main() {
	SetupCb()

	var i int64 = 0
	for {
		if i%1000 == 0 {
			time.Sleep(3 * time.Second)
		}
		go func(i int64) {

			err := hystrix.Do("my_command", func() error {
				// talk to other services
				time.Sleep(1 * time.Second)

				if i%1000 == 0 {
					return errors.New("odd num")
				}

				logrus.Info("success")
				return nil
			}, nil)

			if err != nil {
				logrus.Error(err)
			}
		}(i)
		i++
		if i == 10000 {
			break
		}
	}
}
