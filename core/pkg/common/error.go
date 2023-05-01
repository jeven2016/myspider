package common

import "errors"

var InvalidHomeUrlErr = errors.New("the home url is not set")
var JsonConvertErr = errors.New("unable convert data into json")
