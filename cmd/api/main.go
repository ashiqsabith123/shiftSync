package main

import "shiftsync/pkg/di"

func main() {
	e := di.InitializeAPI()
	e.Start()
}
