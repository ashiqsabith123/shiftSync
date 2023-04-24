package main

import (
	"fmt"
	"shiftsync/pkg/config"
	"shiftsync/pkg/di"
)

func main() {

	config, err := config.LoadConfig()

	e := di.InitializeAPI(config)
	e.Start()
	fmt.Println(err, config)
}
