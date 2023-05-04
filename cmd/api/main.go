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
	server := di.InitializeAPI(config)
	server.Start()
	fmt.Println(err, config)
}
