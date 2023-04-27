package main

import (
	"fmt"
	"shiftsync/pkg/config"
	"shiftsync/pkg/di"
	"shiftsync/pkg/verification"
)

func main() {
	config, err := config.LoadConfig()
	verification.InitTwilio(config)
	e := di.InitializeAPI(config)
	e.Start()
	fmt.Println(err, config)
}
