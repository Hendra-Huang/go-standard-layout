package env

import (
	"testing"

	"github.com/tokopedia/vehicle-insurance/backend/testingutil"
)

func TestGetDefaultValue(t *testing.T) {
	e := Get()
	testingutil.Equals(t, Development, e)
}
