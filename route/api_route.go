package route

import (
	"github.com/dappsteros-io/DappsterOS-MessageBus/codegen"
	"github.com/dappsteros-io/DappsterOS-MessageBus/service"
	jsoniter "github.com/json-iterator/go"
)

type APIRoute struct {
	services *service.Services
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewAPIRoute(services *service.Services) codegen.ServerInterface {
	return &APIRoute{
		services: services,
	}
}
