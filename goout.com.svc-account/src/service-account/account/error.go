package account

import (
	"error"
)

var ErrAccountNull = errors.New("account null")
var ErrPasswordNull = errors.New("password null")
var ErrPasswordNotEnoughLen = errors.New("password not enough length")
var ErrUnknow = errors.New("unknow")
