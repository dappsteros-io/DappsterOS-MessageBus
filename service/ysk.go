package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dappsteros-io/DappsterOS-Common/utils/logger"
	"github.com/dappsteros-io/DappsterOS-MessageBus/common"
	"github.com/dappsteros-io/DappsterOS-MessageBus/model"
	"github.com/dappsteros-io/DappsterOS-MessageBus/pkg/ysk"
	"github.com/dappsteros-io/DappsterOS-MessageBus/repository"
	"github.com/dappsteros-io/DappsterOS-MessageBus/utils"
	"go.uber.org/zap"
)

type YSKService struct {
	repository       *repository.Repository
	ws               *EventServiceWS
	eventTypeService *EventTypeService
}

func NewYSKService(
	repository *repository.Repository,
	ws *EventServiceWS,
	ets *EventTypeService,
) *YSKService {
	return &YSKService{
		repository:       repository,
		ws:               ws,
		eventTypeService: ets,
	}
}

func (s *YSKService) YskCardList(ctx context.Context) ([]ysk.YSKCard, error) {
	cardList, err := (*s.repository).GetYSKCardList()
	if err != nil {
		return []ysk.YSKCard{}, err
	}
	return cardList, nil
}

func (s *YSKService) UpsertYSKCard(ctx context.Context, yskCard ysk.YSKCard) error {
	// don't store short notice cards
	if yskCard.CardType == ysk.CardTypeShortNote {
		return nil
	}
	err := (*s.repository).UpsertYSKCard(yskCard)
	return err
}

func (s *YSKService) DeleteYSKCard(ctx context.Context, id string) error {
	return (*s.repository).DeleteYSKCard(id)
}

func (s *YSKService) Start(init bool) {
	if init {

		s.UpsertYSKCard(context.Background(), utils.ZimaOSDataStationNotice)
		s.UpsertYSKCard(context.Background(), utils.ZimaOSFileManagementNotice)
		s.UpsertYSKCard(context.Background(), utils.ZimaOSRemoteAccessNotice)
	}
	// register event
	s.eventTypeService.RegisterEventType(model.EventType{
		SourceID: common.SERVICENAME,
		Name:     common.EventTypeYSKCardUpsert.Name,
	})

	s.eventTypeService.RegisterEventType(model.EventType{
		SourceID: common.SERVICENAME,
		Name:     common.EventTypeYSKCardDelete.Name,
	})

	channel, err := s.ws.Subscribe(common.SERVICENAME, []string{
		common.EventTypeYSKCardUpsert.Name, common.EventTypeYSKCardDelete.Name,
	})
	if err != nil {
		logger.Error("failed to subscribe to event", zap.Error(err))
		return
	}

	go func() {
		for {
			select {
			case event, ok := <-channel:
				if !ok {
					log.Println("channel closed")
				}
				switch event.Name {
				case common.EventTypeYSKCardUpsert.Name:
					var card ysk.YSKCard
					err := json.Unmarshal([]byte(event.Properties[common.PropertyTypeCardBody.Name]), &card)
					if err != nil {
						logger.Error("failed to umarshal ysk card", zap.Error(err))
						continue
					}
					err = s.UpsertYSKCard(context.Background(), card)
					if err != nil {
						logger.Error("failed to upsert ysk card", zap.Error(err))
					}
				case common.EventTypeYSKCardDelete.Name:
					err = s.DeleteYSKCard(context.Background(), event.Properties[common.PropertyTypeCardID.Name])
					if err != nil {
						logger.Error("failed to delete ysk card", zap.Error(err))
					}
				default:
					fmt.Println(event)
				}
			}
		}
	}()
}
