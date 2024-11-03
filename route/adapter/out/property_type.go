package out

import (
	"github.com/dappster-io/DappsterOS-MessageBus/codegen"
	"github.com/dappster-io/DappsterOS-MessageBus/model"
)

func PropertyTypeAdapter(propertyType model.PropertyType) codegen.PropertyType {
	return codegen.PropertyType{
		Name: propertyType.Name,
	}
}
