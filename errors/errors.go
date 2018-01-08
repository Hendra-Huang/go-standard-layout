package errors

// errors package inspired and a subset copy of upspin project

import (
	"errors"
	"fmt"
	"runtime"

	"log"
)

type Fields map[string]interface{}

// Errs struct
type Errs struct {
	err error

	// Traces used to add function traces to errors, this is different from context
	// While context is used to add more information about the error, traces is used
	// for easier function tracing purposes without hurting heap too much
	traces []string

	// Fields is a fields context similar to logrus.Fields
	// Can be used for adding more context to the errors
	fields Fields
}

/*
Errs will parse arguments based on the data type
1. If string then it will convert the arg to Error
2. If error, then it will just copy the error
3. If the type is *Errs, it will copy the address and create new Errs object
4. If the type is Codes or uint8, then it will convert it to code
*/

// New Errs
func New(args ...interface{}) *Errs {
	var (
		er     error
		traces []string
	)
	err := &Errs{}
	for _, arg := range args {
		switch arg.(type) {
		case string:
			er = errors.New(arg.(string))
		case *Errs:
			// copy and put the errors back
			err := *arg.(*Errs)
			er = err.err
			traces = err.traces
		case error:
			er = arg.(error)
		case Fields:
			if er == nil {
				er = errors.New("error not defined")
			}
			err.fields = arg.(Fields)
		default:
			_, file, line, _ := runtime.Caller(1)
			log.Printf("errors.Errs: bad call from %s:%d: %v", file, line, args)
			er = errors.New("unknown error")
		}
	}
	err.err = er
	err.traces = traces
	return err
}

func (e *Errs) Error() string {
	errorMessage := e.err.Error()
	if len(e.fields) > 0 {
		fieldsContext := ""
		fields := ""
		for key, field := range e.fields {
			fields += fmt.Sprintf("%s:%v, ", key, field)
		}
		fieldsContext = "{ fields: { " + fields[:len(fields)-2] + " } }"
		errorMessage += " " + fieldsContext
	}
	if len(e.traces) > 0 {
		tracesContext := ""
		traces := ""
		for _, trace := range e.traces {
			traces += trace + ", "
		}
		tracesContext = "{ traces: [ " + traces[:len(traces)-2] + " ] }"
		errorMessage += " " + tracesContext
	}

	return errorMessage
}

// GetTraces return traces
func (e *Errs) GetTraces() []string {
	return e.traces
}

func (e *Errs) GetFields() Fields {
	return e.fields
}
