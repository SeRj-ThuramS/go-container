package container

import (
	"errors"
	"reflect"
)

type RegisterStatus int

const (
	RegisterUnknown RegisterStatus = iota + 0
	RegisterCreate
	RegisterCreateInit
	RegisterExists
)

var (
	providers map[any]any
)

func typeProviderReader[T any]() string {
	return reflect.TypeOf((*T)(nil)).String()
}

// RegisterRef function
func RegisterRef[T any](initType ...*T) (RegisterStatus, error) {
	if providers == nil {
		return RegisterUnknown, errors.New("the module's init function was not called")
	}

	typeProvider := typeProviderReader[T]()

	// check is created
	if _, ok := providers[typeProvider]; ok {
		return RegisterExists, nil
	}

	// create new
	if len(initType) == 0 {
		providers[typeProvider] = new(T)

		return RegisterCreate, nil
	}

	if typeProvider != reflect.ValueOf(initType[0]).Type().String() {
		return RegisterUnknown, errors.New("initialization parameter not specified")
	}

	providers[typeProvider] = initType[0]

	return RegisterCreateInit, nil
}

func AssignRef[T any]() (*T, error) {
	typeProvider := typeProviderReader[T]()

	if _, ok := providers[typeProvider]; !ok {
		return nil, errors.New("container <" + typeProvider + "> not found")
	}

	provider, ok := providers[typeProvider].(*T)
	if !ok {
		return nil, errors.New("type provider <" + typeProvider + "> is not registered")
	}

	return provider, nil
}

// init  function
func init() {
	if providers != nil {
		return
	}

	providers = make(map[any]any)
}
