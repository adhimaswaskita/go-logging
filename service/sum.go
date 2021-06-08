package service

import (
	"errors"
	"github.com/adhimaswaskita/go-logging/util"
)

//Divide return number from a value divided by b value and error if any
func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New(util.ErrDividedByZero)
	}

	return a / b, nil
}