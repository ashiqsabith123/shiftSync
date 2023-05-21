package main

import (
	"fmt"
	"shiftsync/notification"
	"shiftsync/pkg/config"
	"shiftsync/pkg/di"
	"shiftsync/pkg/verification"
)

func main() {
	config, err := config.LoadConfig()
	verification.InitTwilio(config)
	server := di.InitializeAPI(config)
	go notification.SendNotification()
	server.Start()
	fmt.Println(err, config)
}
