package string

//go:generate counterfeiter -o ./fakeCoercer.go --fake-name fakeCoercer ./ coercer

import (
	"fmt"
	"github.com/golang-interfaces/encoding-ijson"
	"github.com/golang-interfaces/iioutil"
	"github.com/golang-interfaces/ios"
	"github.com/opspec-io/sdk-golang/model"
	"strconv"
)

type coercer interface {
	// Coerce attempts to coerce Value to a String
	Coerce(
		value *model.Value,
	) (string, error)
}

func newCoercer() coercer {
	return _coercer{
		json:   ijson.New(),
		os:     ios.New(),
		ioUtil: iioutil.New(),
	}
}

type _coercer struct {
	ioUtil iioutil.IIOUtil
	json   ijson.IJSON
	os     ios.IOS
}

func (c _coercer) Coerce(
	value *model.Value,
) (string, error) {
	switch {
	case nil == value:
		return "", nil
	case nil != value.Dir:
		return "", fmt.Errorf("Unable to coerce dir '%v' to string; incompatible types", *value.Dir)
	case nil != value.File:
		fileBytes, err := c.ioUtil.ReadFile(*value.File)
		if nil != err {
			return "", fmt.Errorf("Unable to coerce file to string; error was %v", err.Error())
		}
		return string(fileBytes), nil
	case nil != value.Number:
		return strconv.FormatFloat(*value.Number, 'f', -1, 64), nil
	case nil != value.Object:
		valueBytes, err := c.json.Marshal(value.Object)
		if nil != err {
			return "", fmt.Errorf("Unable to coerce object to string; error was %v", err.Error())
		}
		return string(valueBytes), nil
	case nil != value.String:
		return *value.String, nil
	default:
		return "", fmt.Errorf("Unable to coerce '%v' to string", value)
	}
}