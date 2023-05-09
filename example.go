package main

import (
	"fmt"
	"log"

	"github.com/ErSauravAdhikari/GoCaproverAPI/crapi"
)

func example() {
	// Create a new Caprover instance
	caprover, err := crapi.NewCaproverInstance("https://your-endpoint.xyz", "your-password")
	if err != nil {
		log.Println(err)
		return
	}

	// Get details of all the apps
	appDetails, err := caprover.GetAppDetails()
	if err != nil {
		log.Fatal(err)
	}

	// Print the app details
	for _, app := range appDetails.Data.AppDefinitions {
		fmt.Println("App Name:", app.AppName)
		fmt.Println("Instance Count:", app.InstanceCount)
		fmt.Println("Ports:", app.Ports)
		fmt.Println()
	}
}
