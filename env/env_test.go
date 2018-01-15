package env_test

import (
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/env"
	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestGetDefaultValue(t *testing.T) {
	e := env.Get()
	testingutil.Equals(t, env.Development, e)
}
