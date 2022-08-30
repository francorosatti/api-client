package form3

type Environment string

const (
	EnvironmentLocal      = "local"
	EnvironmentTest       = "test"
	EnvironmentProduction = "production"

	_apiVersion = "v1"
)

var (
	_hostByEnvironment = map[Environment]string{
		EnvironmentLocal:      "http://localhost:8080",
		EnvironmentTest:       "https://internal.form3.com/test",
		EnvironmentProduction: "https://internal.form3.com",
	}
)
