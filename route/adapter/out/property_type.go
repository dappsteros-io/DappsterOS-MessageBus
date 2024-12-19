package out

import (
	"github.com/dappsteros-io/DappsterOS-MessageBus/codegen"
	"github.com/dappsteros-io/DappsterOS-MessageBus/model"
)

func PropertyTypeAdapter(propertyType model.PropertyType) codegen.PropertyType {
	return codegen.PropertyType{
		Name: propertyType.Name,
	}
}
