package env

import (
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestGetDefaultValue(t *testing.T) {
	e := Get()
	testingutil.Equals(t, Development, e)
}
