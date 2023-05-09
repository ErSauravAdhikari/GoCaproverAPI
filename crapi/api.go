package crapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Caprover struct {
	Endpoint string
	Password string
	Token    string
}

// NewCaproverInstance (endpoint string, password string) (Caprover, error): This
// method is a constructor function that creates a new instance of the Caprover
// struct. It takes an endpoint and password as parameters and initializes the
// Caprover struct with the provided values. It also calls the Login method
// internally to authenticate with the Caprover instance using the provided
// credentials.
func NewCaproverInstance(endpoint string, password string) (Caprover, error) {
	cp := Caprover{
		Endpoint: endpoint,
		Password: password,
		Token:    "",
	}

	err := cp.Login()
	if err != nil {
		return Caprover{}, err
	}

	return cp, nil
}

func (c *Caprover) buildURL(path string) string {
	return c.Endpoint + path
}

func (c *Caprover) addHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("x-namespace", "captain")

	if c.Token != "" {
		req.Header.Add("x-captain-auth", c.Token)
	}
}

// Login () error: This method authenticates the client with the Caprover
// instance. It sends a POST request to the Caprover login endpoint with the
// provided password. If the login is successful, it retrieves and stores the
// authentication token for subsequent requests.
func (c *Caprover) Login() error {
	fmt.Println("Attempting Login to Caprover Instance")

	url := c.buildURL(URLLoginPath)

	data := make(map[string]string)
	data["password"] = c.Password
	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("POST", url, payload)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("login Error")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var rsp LoginResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	c.Token = rsp.Data.Token
	return nil
}

// GetAppDetails () (ListAppResponse, error): This method retrieves the details
// of all the applications deployed on the Caprover instance. It sends a GET
// request to the Caprover app list endpoint and returns the list of applications
// along with their details.
func (c *Caprover) GetAppDetails() (ListAppResponse, error) {
	fmt.Println("Getting App Details")

	url := c.buildURL(URLAppListPath)

	data := make(map[string]string)
	data["password"] = c.Password
	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("GET", url, payload)
	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ListAppResponse{}, err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var rsp ListAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	return rsp, nil
}

// GetAppDetailFor (appName string) (AppDefinition, error): This method retrieves the details of
// a specific application by its name. It calls the GetAppDetails method
// internally to get the list of all applications and then searches for the
// application with the matching name. If found, it returns the application
// details; otherwise, it returns an error.
func (c *Caprover) GetAppDetailFor(appName string) (AppDefinition, error) {
	allDetails, _ := c.GetAppDetails()
	for _, v := range allDetails.Data.AppDefinitions {
		if strings.Compare(appName, v.AppName) == 0 {
			return v, nil
		}
	}
	return AppDefinition{}, errors.New("not found")
}

// GetDefaultUpdateRequest (appName string) (UpdateAppRequest, error): This
// method retrieves the default update request for a specific application. It
// calls the GetAppDetails method internally to get the list of all applications
// and then searches for the application with the matching name. If found, it
// returns an UpdateAppRequest containing the default values for updating the
// application; otherwise, it returns an error.
func (c *Caprover) GetDefaultUpdateRequest(appName string) (UpdateAppRequest, error) {
	allDetails, _ := c.GetAppDetails()

	var m AppDefinition
	var found bool
	for _, v := range allDetails.Data.AppDefinitions {
		if strings.Compare(appName, v.AppName) == 0 {
			m = v
			found = true
			break
		}
	}

	if !found {
		return UpdateAppRequest{}, errors.New("not found")
	}

	appRequest := UpdateAppRequest{
		AppName:                           m.AppName,
		InstanceCount:                     m.InstanceCount,
		CaptainDefinitionRelativeFilePath: m.CaptainDefinitionRelativeFilePath,
		NotExposeAsWebApp:                 m.NotExposeAsWebApp,
		ForceSsl:                          m.ForceSsl,
		WebsocketSupport:                  m.WebsocketSupport,
		Volumes:                           m.Volumes,
		Ports:                             m.Ports,
		AppPushWebhook: AppPushWebHook{
			RepoInfo: m.AppPushWebhook.RepoInfo,
		},
		NodeID:                m.NodeID,
		PreDeployFunction:     m.PreDeployFunction,
		ServiceUpdateOverride: m.ServiceUpdateOverride,
		ContainerHTTPPort:     m.ContainerHTTPPort,
		Description:           m.Description,
		EnvVars:               m.EnvVars,
		AppDeployTokenConfig:  m.AppDeployTokenConfig,
	}

	return appRequest, nil
}

// CreateApp (appName string, hasPersistentData bool) error: This method creates
// a new application on the Caprover instance. It sends a POST request to the
// Caprover app register endpoint with the provided appName and hasPersistentData
// parameters. If the creation is successful, it returns nil; otherwise, it
// returns an error.
func (c *Caprover) CreateApp(appName string, hasPersistentData bool) error {
	fmt.Println("Attempting to create a new app")

	url := c.buildURL(URLAppRegisterPath)

	data := make(map[string]interface{})
	data["appName"] = appName
	data["hasPersistentData"] = hasPersistentData

	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("POST", url, payload)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var rsp GenericAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == 100 {
		return nil
	}

	return errors.New(rsp.Description)
}

// updateAppDetails (data UpdateAppRequest) error: This method updates the
// details of an application on the Caprover instance. It sends a POST request to
// the Caprover app update endpoint with the provided UpdateAppRequest payload.
// If the update is successful, it returns nil; otherwise, it returns an error.
// FOR INTERNAL USE ONLY
func (c *Caprover) updateAppDetails(data UpdateAppRequest) error {
	fmt.Println("Attempting to Update App Details")

	url := c.buildURL(URLUpdateAppPath)

	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("POST", url, payload)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var rsp GenericAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == 100 {
		return nil
	}

	return errors.New(rsp.Description)
}

// ForceBuild (token string) error: This method triggers a forced build for an
// application on the Caprover instance. It sends a POST request to the Caprover
// app trigger build endpoint with the provided token parameter. If the build is
// successful, it returns nil; otherwise, it returns an error.
func (c *Caprover) ForceBuild(token string) error {
	fmt.Println("Attempting to Force Build")

	url := c.buildURL(URLAppTriggerBuild) + "?namespace=captain&token=" + token

	req, _ := http.NewRequest("POST", url, nil)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var rsp GenericAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == 100 {
		return nil
	}

	return errors.New(rsp.Description)

}

// EnableBaseDomainSSL (appName string) error: This method enables SSL on the
// base domain for an application. It sends a POST request to the Caprover enable
// base domain SSL endpoint with the provided appName parameter. If the SSL
// enablement is successful, it returns nil; otherwise, it returns an error.
func (c *Caprover) EnableBaseDomainSSL(appName string) error {
	fmt.Println("Attempting to Enable SSL on Base Domain")

	url := c.buildURL(URLEnableBaseDomainSslPath)

	data := make(map[string]string)
	data["appName"] = appName
	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("POST", url, payload)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var rsp GenericAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == 100 {
		return nil
	}

	return errors.New(rsp.Description)
}

// AddCustomDomain (appName string, domain string) error: This method adds a
// custom domain to an application. It sends a POST request to the Caprover add
// custom domain endpoint with the provided appName and domain parameters. If the
// domain addition is successful, it returns nil; otherwise, it returns an error.
func (c *Caprover) AddCustomDomain(appName string, domain string) error {
	fmt.Println("Attempting to add a new domain")

	url := c.buildURL(URLAddCustomDomainPath)

	data := make(map[string]string)
	data["appName"] = appName
	data["customDomain"] = domain
	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("POST", url, payload)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var rsp GenericAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == 100 {
		return nil
	}

	return errors.New(rsp.Description)
}

// EnableCustomDomainSSL (appName string, domain string) error: This method
// enables SSL on a custom domain for an application. It sends a POST request to
// the Caprover enable custom domain SSL endpoint with the provided appName and
// domain parameters. If the SSL enablement is successful, it returns nil;
// otherwise, it returns an error.
func (c *Caprover) EnableCustomDomainSSL(appName string, domain string) error {
	fmt.Println("Attempting to Enable SSL on Custom Domain")

	url := c.buildURL(URLEnableCustomDomainSslPath)

	data := make(map[string]string)
	data["appName"] = appName
	data["customDomain"] = domain
	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("POST", url, payload)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var rsp GenericAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == 100 {
		return nil
	}

	return errors.New(rsp.Description)
}

// RestartApp restarts app with given appName
func (c *Caprover) RestartApp(appName string) error {
	err := c.updateAppDetails(UpdateAppRequest{
		AppName: appName,
	})

	return err
}

func (c *Caprover) UpdateContainerHTTPPort(appName string, newPort int) error {
	err := c.updateAppDetails(UpdateAppRequest{
		AppName:           appName,
		ContainerHTTPPort: newPort,
	})

	return err
}

func (c *Caprover) EnableWebsocketSupport(appName string) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	currentConfig.WebsocketSupport = true

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) EnableForceHTTPS(appName string) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	currentConfig.ForceSsl = true

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) DisableWebsocketSupport(appName string) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	currentConfig.WebsocketSupport = false

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) DisableForceHTTPS(appName string) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	currentConfig.ForceSsl = false

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) TurnInstanceCountZero(appName string) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	currentConfig.InstanceCount = 0

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) TurnInstanceCountOne(appName string) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	currentConfig.InstanceCount = 1

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) UpdateGitRepoInfo(appName string, repoInfo AppRepoInfo) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	currentConfig.AppPushWebhook.RepoInfo = repoInfo

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) UpdateResourceConstraint(appName string, memoryInMB int64, cpuInUnits float64) error {
	currentConfig, err := c.GetDefaultUpdateRequest(appName)

	suo, err := json.Marshal(ServiceUpdateOverride{
		TaskTemplate: SUOTaskTemplate{
			Resources: SUOResources{
				Limits: SUOLimits{
					MemoryBytes: memoryInMB * ResourceOneMb,
					NanoCPUs:    int64(cpuInUnits * float64(ResourceOneCpu)),
				},
			},
		},
	})

	currentConfig.ServiceUpdateOverride = string(suo)

	if err != nil {
		return err
	}

	err = c.updateAppDetails(currentConfig)

	return err
}

func (c *Caprover) GetBuildLogs(appName string) (string, error) {
	fmt.Println("Getting Build Logs")

	url := c.buildURL(URLAppBuildLog) + "/" + appName + "/"

	req, _ := http.NewRequest("GET", url, nil)
	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var rsp AppBuildLogResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	logLines := strings.Join(rsp.Data.Logs.Lines, "\n")

	return logLines, nil
}

func (c *Caprover) GetAppLogs(appName string) (string, error) {
	fmt.Println("Getting App Logs")

	url := c.buildURL(URLAppBuildLog) + "/" + appName + "/logs"

	req, _ := http.NewRequest("GET", url, nil)
	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var rsp AppLogResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	logLines := rsp.Data.Logs

	return logLines, nil
}

// RemoveApp (appName string) error`: This method deletes an application from the
// Caprover instance. It deletes a given Caprover app based on the provided
// `appName` parameter. If the deletion is successful, it returns nil; otherwise,
// it returns an error.
func (c *Caprover) RemoveApp(appName string) error {
	fmt.Println("Attempting to Remove an APP")

	url := c.buildURL(URLAppDeletePath)

	data := make(map[string]string)
	data["appName"] = appName
	jsonEncode, _ := json.Marshal(data)
	payload := bytes.NewBuffer(jsonEncode)

	req, _ := http.NewRequest("POST", url, payload)

	c.addHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var rsp GenericAppResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		log.Fatal(err)
	}

	if rsp.Status == 100 {
		return nil
	}

	return errors.New(rsp.Description)
}
