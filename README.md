# GoCaproverAPI

GoCaproverAPI is a pure Go interface for managing deployments in Caprover, a deployment management software. It provides an easy-to-use API for interacting with Caprover and performing various deployment-related operations.

## Installation

To use GoCaproverAPI in your Go project, you can use the following command to add it as a dependency:

```shell
go get github.com/ErSauravAdhikari/GoCaproverAPI
```

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

- `NewCaproverInstance(endpoint string, password string) (Caprover, error)`: Creates a new Caprover instance and performs login.
- `GetAppDetails() (ListAppResponse, error)`: Retrieves details of all the apps.
- `GetAppDetailFor(appName string) (AppDefinition, error)`: Retrieves details of a specific app.
- `GetDefaultUpdateRequest(appName string) (UpdateAppRequest, error)`: Retrieves the default update request for an app.
- `CreateApp(appName string, hasPersistentData bool) error`: Creates a new app.
- `UpdateContainerHTTPPort(appName string, newPort int) error`: Updates the container HTTP port for an app.
- `RestartApp(appName string) error`: Restarts an app.
- `EnableWebsocketSupport(appName string) error`: Enables websocket support for an app.
- `DisableWebsocketSupport(appName string) error`: Disables websocket support for an app.
- `EnableForceHTTPS(appName string) error`: Enables force HTTPS for an app.
- `DisableForceHTTPS(appName string) error`: Disables force HTTPS for an app.
- `TurnInstanceCountZero(appName string) error`: Sets the instance count of an app to zero.
- `TurnInstanceCountOne(appName string) error`: Sets the instance count of an app to one.
- `UpdateGitRepoInfo(appName string, repoInfo AppRepoInfo) error`: Updates the Git repository info for an app.
- `UpdateResourceConstraint(appName string, memoryInMB int64, cpuInUnits float64) error`: Updates the resource constraints for an app.
- `GetBuildLogs(appName string) (string, error)`: Retrieves the build logs for an app.
- `GetAppLogs(appName string) (string, error)`: Retrieves the app logs for an app.
- `RemoveApp(appName string) error`: Removes an app.

Please refer to the code comments and the GoCaproverAPI package documentation for more details on each method
