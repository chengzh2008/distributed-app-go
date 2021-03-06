package registry

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	RegService     = ServiceName("RegistryService")
	GradingService = ServiceName("GradingService")
)
