package main

import (
	"github.com/Hendra-Huang/go-standard-layout/log"
)

type errors struct {
	fields map[string]interface{}
}

func (e errors) GetFields() map[string]interface{} {
	return e.fields
}

func (e errors) Error() string {
	return "Error sample"
}

func main() {
	logDebug()
	logInfo()
	logErrors()
}

func logDebug() {
	// anything above debug will come out
	log.SetLevel(log.DebugLevel)
	log.Debug("This is debug log")
	log.Info("This is info log")
	log.Warn("This is warn log")
	log.Error("This is error log")
}

func logInfo() {
	// anything below info will not come out
	log.SetLevel(log.InfoLevel)
	log.WithFields(log.Fields{"field1": "value1"}).Info("This is info log")
}

func logErrors() {
	fields := map[string]interface{}{"label1": "val1", "label2": "val2"}
	err := errors{fields}
	log.Errors(err)
}
