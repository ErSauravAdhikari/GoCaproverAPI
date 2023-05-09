package crapi

import "time"

// LoginResponse holds the response from the login endpoint. It contains the token information.
type LoginResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Token string `json:"token"`
	} `json:"data"`
}

// VolumeInformation holds a single persistant directory info for a given app.
type VolumeInformation struct {
	ContainerPath string `json:"containerPath"`
	VolumeName    string `json:"volumeName"`
}

// PortInformation holds a single port mapping info for a given app.
type PortInformation struct {
	HostPort      int `json:"hostPort"`
	ContainerPort int `json:"containerPort"`
}

// AppPushWebHook acts as bucket to store the repository information.
type AppPushWebHook struct {
	RepoInfo AppRepoInfo `json:"repoInfo"`
}

// AppRepoInfo acts as bucket to store the repository details along with auth credentials.
type AppRepoInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Branch   string `json:"branch"`
	SSHKey   string `json:"sshKey"`
	Repo     string `json:"repo"`
}

// EnvVarInformation holds a single environment variable information for a given app.
type EnvVarInformation struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AppDeployTokenConfig struct {
	Enabled bool `json:"enabled"`
}

// AppDefinition holds all the information stored by the caprover for a given app.
type AppDefinition struct {
	HasPersistentData                 bool                `json:"hasPersistentData"`
	Description                       string              `json:"description"`
	InstanceCount                     int                 `json:"instanceCount"`
	CaptainDefinitionRelativeFilePath string              `json:"captainDefinitionRelativeFilePath"`
	Networks                          []string            `json:"networks"`
	EnvVars                           []EnvVarInformation `json:"envVars"`
	Volumes                           []VolumeInformation `json:"volumes"`
	Ports                             []PortInformation   `json:"ports"`
	Versions                          []struct {
		Version           int       `json:"version"`
		TimeStamp         time.Time `json:"timeStamp"`
		DeployedImageName string    `json:"deployedImageName"`
		GitHash           string    `json:"gitHash"`
	} `json:"versions"`
	DeployedVersion        int                  `json:"deployedVersion"`
	NotExposeAsWebApp      bool                 `json:"notExposeAsWebApp"`
	CustomDomain           []any                `json:"customDomain"`
	HasDefaultSubDomainSsl bool                 `json:"hasDefaultSubDomainSsl"`
	ForceSsl               bool                 `json:"forceSsl"`
	WebsocketSupport       bool                 `json:"websocketSupport"`
	ContainerHTTPPort      int                  `json:"containerHttpPort"`
	NodeID                 string               `json:"nodeId,omitempty"`
	PreDeployFunction      string               `json:"preDeployFunction"`
	ServiceUpdateOverride  string               `json:"serviceUpdateOverride"`
	AppDeployTokenConfig   AppDeployTokenConfig `json:"appDeployTokenConfig"`
	AppName                string               `json:"appName"`
	IsAppBuilding          bool                 `json:"isAppBuilding"`
	AppPushWebhook         struct {
		TokenVersion     string      `json:"tokenVersion"`
		PushWebhookToken string      `json:"pushWebhookToken"`
		RepoInfo         AppRepoInfo `json:"repoInfo"`
	} `json:"appPushWebhook,omitempty"`
}

// ListAppResponse holds the response for the list all app request.
type ListAppResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Data        struct {
		AppDefinitions     []AppDefinition `json:"appDefinitions"`
		RootDomain         string          `json:"rootDomain"`
		DefaultNginxConfig string          `json:"defaultNginxConfig"`
	} `json:"data"`
}

// GenericAppResponse holds the response for the generic requests. These include delete, etc. requests that don't
// have much data information.
type GenericAppResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Data        any    `json:"data"`
}

// UpdateAppRequest holds response for update app request
type UpdateAppRequest struct {
	AppName                           string               `json:"appName"`
	InstanceCount                     int                  `json:"instanceCount"`
	CaptainDefinitionRelativeFilePath string               `json:"captainDefinitionRelativeFilePath"`
	NotExposeAsWebApp                 bool                 `json:"notExposeAsWebApp"`
	ForceSsl                          bool                 `json:"forceSsl"`
	WebsocketSupport                  bool                 `json:"websocketSupport"`
	Volumes                           []VolumeInformation  `json:"volumes"`
	Ports                             []PortInformation    `json:"ports"`
	AppPushWebhook                    AppPushWebHook       `json:"appPushWebhook"`
	NodeID                            string               `json:"nodeId"`
	PreDeployFunction                 string               `json:"preDeployFunction"`
	ServiceUpdateOverride             string               `json:"serviceUpdateOverride"`
	ContainerHTTPPort                 int                  `json:"containerHttpPort"`
	Description                       string               `json:"description"`
	EnvVars                           []EnvVarInformation  `json:"envVars"`
	AppDeployTokenConfig              AppDeployTokenConfig `json:"appDeployTokenConfig"`
}

// CustomAppRepositoryConfig holds custom app repository information.
type CustomAppRepositoryConfig struct {
	Repository string
	Branch     string
	Username   string
	Password   string
}

// SUOLimits is used to enforce resource constraints on given apps.
type SUOLimits struct {
	MemoryBytes int64 `json:"MemoryBytes"`
	NanoCPUs    int64 `json:"NanoCPUs"`
}

type SUOResources struct {
	Limits SUOLimits `json:"Limits"`
}

type SUOTaskTemplate struct {
	Resources SUOResources `json:"Resources"`
}

// ServiceUpdateOverride is a bucket that holds information used to enforce resource constraints on given apps.
type ServiceUpdateOverride struct {
	TaskTemplate SUOTaskTemplate `json:"TaskTemplate"`
}

// AppBuildLogLogs stores the actual build logs as returned by the api
type AppBuildLogLogs struct {
	Lines []string `json:"lines"`
}

// AppBuildLogData is a data bucket for AppBuildLogLogs
type AppBuildLogData struct {
	Logs AppBuildLogLogs `json:"logs"`
}

// AppBuildLogResponse is a response bucket for AppBuildLogData
type AppBuildLogResponse struct {
	Status      int             `json:"status"`
	Description string          `json:"description"`
	Data        AppBuildLogData `json:"data"`
}

// AppLogData stores the actual app logs as returned by the api
type AppLogData struct {
	Logs string `json:"logs"`
}

// AppLogResponse is a response bucket for AppLogData
type AppLogResponse struct {
	Status      int        `json:"status"`
	Description string     `json:"description"`
	Data        AppLogData `json:"data"`
}
