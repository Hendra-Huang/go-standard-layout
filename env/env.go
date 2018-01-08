package env

/*
This env package is used for this project only
Will return all defined env value from TKPENV
*/

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

	envVar = "TKPENV"
)

func Get() Environment {
	e := os.Getenv(envVar)
	if e == "" {
		return Development
	}
	return Environment(e)
}
