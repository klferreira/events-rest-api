package event

import (
	"context"

	"github.com/klferreira/events-rest-api/pkg/model"
)

type Service interface {
	Fetch(ctx context.Context, filters interface{}) ([]*model.Event, error)
	Create(ctx context.Context, event *model.Event) (*model.Event, error)
	Update(ctx context.Context, event *model.Event) (*model.Event, error)
	Delete(ctx context.Context, event *model.Event) (*model.Event, error)
}