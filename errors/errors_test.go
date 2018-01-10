package errors

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Hendra-Huang/go-standard-layout/testingutil"
)

func TestErrs(t *testing.T) {
	cases := []struct {
		err    *Errs
		expect *Errs
	}{
		{
			err: New("New error without anything"),
			expect: &Errs{
				err: errors.New("New error without anything"),
			},
		},
		{
			err: New("Error with fields", Fields{"first": "two", "satu": 2}),
			expect: &Errs{
				err:    errors.New("Error with fields"),
				fields: Fields{"first": "two", "satu": 2},
			},
		},
	}

	for _, val := range cases {
		if !reflect.DeepEqual(val.err, val.expect) {
			testingutil.Equals(t, val.expect, val.err)
		}
	}
}

func TestError(t *testing.T) {
	cases := []struct {
		input  *Errs
		expect string
	}{
		{
			input:  New("Input as string"),
			expect: "Input as string",
		},
		{
			input:  New(errors.New("Input as string")),
			expect: "Input as string",
		},
		{
			input:  New("Error with fields", Fields{"a": "abc"}),
			expect: "Error with fields { fields: { a:abc } }",
		},
	}

	for _, val := range cases {
		testingutil.Equals(t, val.expect, val.input.Error())
	}
}
