package graph

import (
	"myapp/service"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var Broker *service.Broker

type Resolver struct{}
