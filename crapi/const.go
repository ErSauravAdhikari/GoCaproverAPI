package crapi

const (
	ResourceOneMb  int64 = 1048576
	ResourceOneCpu int64 = 1000000000

	URLLoginPath                 = "/api/v2/login"
	URLAppListPath               = "/api/v2/user/apps/appDefinitions"
	URLAppRegisterPath           = "/api/v2/user/apps/appDefinitions/register"
	URLUpdateAppPath             = "/api/v2/user/apps/appDefinitions/update"
	URLAppTriggerBuild           = "/api/v2/user/apps/webhooks/triggerbuild"
	URLEnableBaseDomainSslPath   = "/api/v2/user/apps/appDefinitions/enablebasedomainssl"
	URLAddCustomDomainPath       = "/api/v2/user/apps/appDefinitions/customdomain"
	URLEnableCustomDomainSslPath = "/api/v2/user/apps/appDefinitions/enablecustomdomainssl"
	URLAppBuildLog               = "/api/v2/user/apps/appData"
	URLAppDeletePath             = "/api/v2/user/apps/appDefinitions/delete"
)
