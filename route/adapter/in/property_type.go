package in

import (
	"github.com/dappster-io/DappsterOS-MessageBus/codegen"
	"github.com/dappster-io/DappsterOS-MessageBus/model"
)

func PropertyTypeAdapter(propertyType codegen.PropertyType) model.PropertyType {
	return model.PropertyType{
		Name: propertyType.Name,
	}
}
