package main

import (
	"fmt"
	"github.com/rafaph/notte-auth/config"
)

func main() {
	conf, err := config.NewConfig()
	fmt.Println(err)
	fmt.Println(conf)
}
