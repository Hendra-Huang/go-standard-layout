package env

import "os"

type Environment string

func (e Environment) String() string {
	return string(e)
}

const (
	Development Environment = "dev"
	Alpha       Environment = "alpha"
	Staging     Environment = "staging"
	Production  Environment = "prod"

	envVar = "MYAPPENV"
)

func App() Environment {
	e := os.Getenv(envVar)
	if e == "" {
		return Development
	}
	return Environment(e)
}

func Get(name string) string {
	return os.Getenv(name)
}

func GetWithDefault(name, defaultValue string) string {
	e := Get(name)
	if e != "" {
		return e
	}
	return defaultValue
}
