package out

import (
	"time"

	"github.com/dappster-io/DappsterOS-Common/utils"
	"github.com/dappster-io/DappsterOS-MessageBus/codegen"
	"github.com/dappster-io/DappsterOS-MessageBus/model"
)

func ActionAdapter(action model.Action) codegen.Action {
	return codegen.Action{
		SourceID:   action.SourceID,
		Name:       action.Name,
		Properties: action.Properties,
		Timestamp:  utils.Ptr(time.Unix(action.Timestamp, 0)),
	}
}
