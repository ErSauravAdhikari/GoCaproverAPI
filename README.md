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

## Detailed Method Details
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

11. `EnableCustomDomainSSL(appName string, domain string) error`: This method enables SSL on a custom domain for an application. It sends a POST request to the Caprover enable custom domain SSL endpoint with the provided