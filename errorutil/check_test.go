package errorutil

import (
	"errors"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestCheckWithErrorHandler(t *testing.T) {
	cases := []struct {
		err    error
		expect bool
	}{
		{
			err:    errors.New("error"),
			expect: true,
		},
		{
			err:    nil,
			expect: false,
		},
	}

	for _, val := range cases {
		var isError bool
		CheckWithErrorHandler(val.err, func(err error) {
			isError = true
		})
		testingutil.Equals(t, val.expect, isError)
	}
}

func TestIsError(t *testing.T) {
	cases := []struct {
		err    error
		expect bool
	}{
		{
			err:    errors.New("error"),
			expect: true,
		},
		{
			err:    nil,
			expect: false,
		},
	}

	for _, val := range cases {
		testingutil.Equals(t, val.expect, IsError(val.err))
	}
}
