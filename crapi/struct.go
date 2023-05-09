package crapi

import "time"

type LoginResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Token string `json:"token"`
	} `json:"data"`
}
type VolumeInformation struct {
	ContainerPath string `json:"containerPath"`
	VolumeName    string `json:"volumeName"`
}

type PortInformation struct {
	HostPort      int `json:"hostPort"`
	ContainerPort int `json:"containerPort"`
}

type AppPushWebHook struct {
	RepoInfo AppRepoInfo `json:"repoInfo"`
}

type AppRepoInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Branch   string `json:"branch"`
	SSHKey   string `json:"sshKey"`
	Repo     string `json:"repo"`
}

type EnvVarInformation struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AppDeployTokenConfig struct {
	Enabled bool `json:"enabled"`
}

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

type ListAppResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Data        struct {
		AppDefinitions     []AppDefinition `json:"appDefinitions"`
		RootDomain         string          `json:"rootDomain"`
		DefaultNginxConfig string          `json:"defaultNginxConfig"`
	} `json:"data"`
}

type GenericAppResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Data        any    `json:"data"`
}

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

type CustomAppRepositoryConfig struct {
	Repository string
	Branch     string
	Username   string
	Password   string
}

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

type ServiceUpdateOverride struct {
	TaskTemplate SUOTaskTemplate `json:"TaskTemplate"`
}

type AppBuildLogLogs struct {
	Lines []string `json:"lines"`
}

type AppBuildLogData struct {
	Logs AppBuildLogLogs `json:"logs"`
}

type AppBuildLogResponse struct {
	Status      int             `json:"status"`
	Description string          `json:"description"`
	Data        AppBuildLogData `json:"data"`
}

type AppLogData struct {
	Logs string `json:"logs"`
}

type AppLogResponse struct {
	Status      int        `json:"status"`
	Description string     `json:"description"`
	Data        AppLogData `json:"data"`
}
