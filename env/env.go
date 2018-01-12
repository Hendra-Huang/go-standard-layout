package env

import "os"

type Environment string

func (e Environment) String() string {
	return string(e)
}

const (
	Development Environment = "development"
	Alpha       Environment = "alpha"
	Staging     Environment = "staging"
	Production  Environment = "production"

	envVar = "MYAPPENV"
)

func Get() Environment {
	e := os.Getenv(envVar)
	if e == "" {
		return Development
	}
	return Environment(e)
}
