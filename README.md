# GoCaproverAPI!

![gophe r](https://github.com/ErSauravAdhikari/GoCaproverAPI/assets/69170305/4a846cbf-a342-4402-96c9-2e513f4b3823)

GoCaproverAPI is a pure Go wrapper for caprover api for managing deployments in Caprover, a self hosted PAAS software. 
It provides an easy-to-use API for interacting with Caprover and performing various deployment-related operations.

This also adds additional functionality such as easy to add resource constraints that is not as easy to add using the official caprover webui.

Caprover Docs: https://caprover.com/

## Installation

To use GoCaproverAPI in your Go project, you can use the following command to add it as a dependency:

```shell
go get github.com/ErSauravAdhikari/GoCaproverAPI
```

## Documentation
[https://pkg.go.dev/github.com/ErSauravAdhikari/GoCaproverAPI/crapi](https://pkg.go.dev/github.com/ErSauravAdhikari/GoCaproverAPI/crapi#Caprover)

## Usage

Here's a simple example that demonstrates how to use GoCaproverAPI:

```go
package main

import (
	"fmt"
	"log"

	"github.com/ErSauravAdhikari/GoCaproverAPI/crapi"
)

func main() {
	// Create a new Caprover instance
	caprover, err := crapi.NewCaproverInstance("http://your-caprover-endpoint", "your-password")
	if err != nil {
		log.Fatal(err)
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
```

## API Documentation

The following methods are available in the Caprover struct:

1. `NewCaproverInstance(endpoint string, password string) (Caprover, error)`: This method is a constructor function that creates a new instance of the `Caprover` struct. It takes an `endpoint` and `password` as parameters and initializes the `Caprover` struct with the provided values. It also calls the `Login` method internally to authenticate with the Caprover instance using the provided credentials.

2. `Login() error`: This method authenticates the client with the Caprover instance. It sends a POST request to the Caprover login endpoint with the provided password. If the login is successful, it retrieves and stores the authentication token for subsequent requests.

3. `GetAppDetails() (ListAppResponse, error)`: This method retrieves the details of all the applications deployed on the Caprover instance. It sends a GET request to the Caprover app list endpoint and returns the list of applications along with their details.

4. `GetAppDetailFor(appName string) (AppDefinition, error)`: This method retrieves the details of a specific application by its name. It calls the `GetAppDetails` method internally to get the list of all applications and then searches for the application with the matching name. If found, it returns the application details; otherwise, it returns an error.

5. `GetDefaultUpdateRequest(appName string) (UpdateAppRequest, error)`: This method retrieves the default update request for a specific application. It calls the `GetAppDetails` method internally to get the list of all applications and then searches for the application with the matching name. If found, it returns an `UpdateAppRequest` containing the default values for updating the application; otherwise, it returns an error.

6. `CreateApp(appName string, hasPersistentData bool) error`: This method creates a new application on the Caprover instance. It sends a POST request to the Caprover app register endpoint with the provided `appName` and `hasPersistentData` parameters. If the creation is successful, it returns nil; otherwise, it returns an error.

7. `updateAppDetails(data UpdateAppRequest) error`: This method updates the details of an application on the Caprover instance. It sends a POST request to the Caprover app update endpoint with the provided `UpdateAppRequest` payload. If the update is successful, it returns nil; otherwise, it returns an error.

8. `ForceBuild(token string) error`: This method triggers a forced build for an application on the Caprover instance. It sends a POST request to the Caprover app trigger build endpoint with the provided `token` parameter. If the build is successful, it returns nil; otherwise, it returns an error.

9. `EnableBaseDomainSSL(appName string) error`: This method enables SSL on the base domain for an application. It sends a POST request to the Caprover enable base domain SSL endpoint with the provided `appName` parameter. If the SSL enablement is successful, it returns nil; otherwise, it returns an error.

10. `AddCustomDomain(appName string, domain string) error`: This method adds a custom domain to an application. It sends a POST request to the Caprover add custom domain endpoint with the provided `appName` and `domain` parameters. If the domain addition is successful, it returns nil; otherwise, it returns an error.

11. `EnableCustomDomainSSL(appName string, domain string) error`: This method enables SSL on a custom domain for an application. It sends a POST request to the Caprover enable custom domain SSL endpoint with the provided appName and domain parameters. If the SSL enablement is successful, it returns nil; otherwise, it returns an error. 

12. `RemoveApp(appName string) error`: This method deletes an application from the Caprover instance. It deletes a given Caprover app based on the provided `appName` parameter. If the deletion is successful, it returns nil; otherwise, it returns an error.


These methods provide a comprehensive set of functionalities to interact with a Caprover instance programmatically. They allow you to perform tasks such as creating and deleting applications, updating application details, scaling instances, managing custom domains, enabling SSL, retrieving and setting persistent data, triggering builds, and deploying applications. By utilizing these methods, you can automate and streamline your interactions with the Caprover platform.

## Example
Sure! Here are the updated examples using the proper structure:

1. Example: Creating a new app with persistent data

```go
func main() {
	// Initialize Caprover instance
	caprover, err := crapi.NewCaproverInstance("https://your-caprover-instance.com", "your-password")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new app with persistent data
	appName := "my-app"
	err = caprover.CreateApp(appName, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created app %s with persistent data.\n", appName)
}
```

2. Example: Enabling SSL on the base domain

```go
func main() {
	// Initialize Caprover instance
	caprover, err := crapi.NewCaproverInstance("https://your-caprover-instance.com", "your-password")
	if err != nil {
		log.Fatal(err)
	}

	// Enable SSL on the base domain for an app
	appName := "my-app"
	err = caprover.EnableBaseDomainSSL(appName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SSL enabled on base domain for app %s.\n", appName)
}
```

3. Example: Adding a custom domain to an app

```go
func main() {
	// Initialize Caprover instance
	caprover, err := crapi.NewCaproverInstance("https://your-caprover-instance.com", "your-password")
	if err != nil {
		log.Fatal(err)
	}

	// Add a custom domain to an app
	appName := "my-app"
	domain := "custom-domain.com"
	err = caprover.AddCustomDomain(appName, domain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added custom domain %s to app %s.\n", domain, appName)
}
```

4. Example: Updating resource constraints for an app

```go
func main() {
	// Initialize Caprover instance
	caprover, err := crapi.NewCaproverInstance("https://your-caprover-instance.com", "your-password")
	if err != nil {
		log.Fatal(err)
	}

	// Update resource constraints for an app
	appName := "my-app"
	memoryInMB := int64(512)
	cpuInUnits := 1.0
	err = caprover.UpdateResourceConstraint(appName, memoryInMB, cpuInUnits)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated resource constraints for app %s.\n", appName)
}
```

Please make sure to replace "https://your-caprover-instance.com" and "your-password" with the actual URL and password of your Caprover instance.
