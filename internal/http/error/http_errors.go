package error

import "errors"

var (
	CepParamNotExists = errors.New("cep parameter can not be blank")
	CepInvalid        = errors.New("invalid zipcode")
	CepCantFind       = errors.New("can not find zipcode")
)
