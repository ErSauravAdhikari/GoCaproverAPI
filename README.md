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

12. `DeleteApp(appName string) error`: This method deletes an application from the Caprover instance. It sends a DELETE request to the Caprover app delete endpoint with the provided `appName` parameter. If the deletion is successful, it returns nil; otherwise, it returns an error.

13. `ScaleApp(appName string, instances int) error`: This method scales the number of instances for an application on the Caprover instance. It sends a POST request to the Caprover app scale endpoint with the provided `appName` and `instances` parameters. If the scaling is successful, it returns nil; otherwise, it returns an error.

14. `GetPersistentData(appName string) (map[string]interface{}, error)`: This method retrieves the persistent data associated with an application on the Caprover instance. It sends a GET request to the Caprover app persistent data endpoint with the provided `appName` parameter. If the retrieval is successful, it returns a map of key-value pairs representing the persistent data; otherwise, it returns an error.

15. `SetPersistentData(appName string, data map[string]interface{}) error`: This method sets the persistent data for an application on the Caprover instance. It sends a POST request to the Caprover app persistent data endpoint with the provided `appName` and `data` parameters. If the data setting is successful, it returns nil; otherwise, it returns an error.

16. `DeployApp(appName string, deploymentData DeploymentData) error`: This method deploys an application on the Caprover instance. It sends a POST request to the Caprover app deploy endpoint with the provided `appName` and `deploymentData` parameters. The `deploymentData` contains information such as the image to deploy, port mapping, environment variables, etc. If the deployment is successful, it returns nil; otherwise, it returns an error.

These methods provide a comprehensive set of functionalities to interact with a Caprover instance programmatically. They allow you to perform tasks such as creating and deleting applications, updating application details, scaling instances, managing custom domains, enabling SSL, retrieving and setting persistent data, triggering builds, and deploying applications. By utilizing these methods, you can automate and streamline your interactions with the Caprover platform.

## Example
1. `ListApps() ([]App, error)`:

```go
apps, err := caprover.ListApps()
if err != nil {
    // handle error
}

for _, app := range apps {
    fmt.Println(app.Name)
}
```

2. `GetApp(appName string) (App, error)`:

```go
app, err := caprover.GetApp("myapp")
if err != nil {
    // handle error
}

fmt.Println("Name:", app.Name)
fmt.Println("Image:", app.Deployment.Image)
```

3. `CreateApp(appData AppData) (string, error)`:

```go
appData := caprover.AppData{
    Name: "myapp",
    Deployment: caprover.DeploymentData{
        Image: "myapp:latest",
        Port: 8080,
    },
}

appID, err := caprover.CreateApp(appData)
if err != nil {
    // handle error
}

fmt.Println("Created app ID:", appID)
```

4. `UpdateApp(appName string, appData AppData) error`:

```go
appData := caprover.AppData{
    Name: "myapp",
    Deployment: caprover.DeploymentData{
        Image: "myapp:v2",
    },
}

err := caprover.UpdateApp("myapp", appData)
if err != nil {
    // handle error
}
```

5. `TriggerBuild(appName string) error`:

```go
err := caprover.TriggerBuild("myapp")
if err != nil {
    // handle error
}
```

6. `GetBuildLogs(appName string) (string, error)`:

```go
logs, err := caprover.GetBuildLogs("myapp")
if err != nil {
    // handle error
}

fmt.Println("Build logs:", logs)
```

7. `SetCustomDomain(appName string, domain string) error`:

```go
err := caprover.SetCustomDomain("myapp", "myapp.example.com")
if err != nil {
    // handle error
}
```

8. `DeleteCustomDomain(appName string) error`:

```go
err := caprover.DeleteCustomDomain("myapp")
if err != nil {
    // handle error
}
```

9. `EnableSSL(appName string, domain string) error`:

```go
err := caprover.EnableSSL("myapp", "myapp.example.com")
if err != nil {
    // handle error
}
```

10. `DisableSSL(appName string) error`:

```go
err := caprover.DisableSSL("myapp")
if err != nil {
    // handle error
}
```

11. `DeleteApp(appName string) error`:

```go
err := caprover.DeleteApp("myapp")
if err != nil {
    // handle error
}
```

12. `ScaleApp(appName string, instances int) error`:

```go
err := caprover.ScaleApp("myapp", 3)
if err != nil {
    // handle error
}
```

13. `GetPersistentData(appName string) (map[string]interface{}, error)`:

```go
data, err := caprover.GetPersistentData("myapp")
if err != nil {
    // handle error
}

fmt.Println("Persistent data:", data)
```

14. `SetPersistentData(appName string, data map[string]interface{}) error`:

Here's the continuation from the `SetPersistentData` function:

```go
data := map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
}

err := caprover.SetPersistentData("myapp", data)
if err != nil {
    // handle error
}
```

15. `DeletePersistentData(appName string) error`:

```go
err := caprover.DeletePersistentData("myapp")
if err != nil {
    // handle error
}
```

16. `GetLogs(appName string) (string, error)`:

```go
logs, err := caprover.GetLogs("myapp")
if err != nil {
    // handle error
}

fmt.Println("Logs:", logs)
```

17. `GetServerLogs() (string, error)`:

```go
logs, err := caprover.GetServerLogs()
if err != nil {
    // handle error
}

fmt.Println("Server logs:", logs)
```

18. `DeployLocalFile(appName string, localFilePath string, remoteFilePath string) error`:

```go
localFilePath := "/path/to/local/file.txt"
remoteFilePath := "/app/data/file.txt"

err := caprover.DeployLocalFile("myapp", localFilePath, remoteFilePath)
if err != nil {
    // handle error
}
```

19. `DeployURL(appName string, url string, remoteFilePath string) error`:

```go
url := "https://example.com/file.txt"
remoteFilePath := "/app/data/file.txt"

err := caprover.DeployURL("myapp", url, remoteFilePath)
if err != nil {
    // handle error
}
```

20. `ExecCommand(appName string, command string) (string, error)`:

```go
command := "ls -l /app/data"

output, err := caprover.ExecCommand("myapp", command)
if err != nil {
    // handle error
}

fmt.Println("Command output:", output)
```

These examples demonstrate the usage of each function in various scenarios with different input parameters. Remember to replace the placeholder values (`myapp`, `localhost`, file paths, etc.) with your specific application or environment details.
