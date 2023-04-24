package main

import (
	"fmt"
	"shiftsync/pkg/config"
)

func main() {
	// e := di.InitializeAPI()
	// e.Start()
	config, err := config.LoadConfig()

	fmt.Println(err, config)
}
