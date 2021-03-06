package event

import (
	"context"

	"github.com/klferreira/events-rest-api/internal/model"
)

type Repository interface {
	Get(ctx context.Context, filters interface{}) ([]*model.Event, error)
	Find(ctx context.Context, id string) (*model.Event, error)
	Create(ctx context.Context, event *model.Event) (*model.Event, error)
	Update(ctx context.Context, event *model.Event) (*model.Event, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type Service interface {
	Fetch(ctx context.Context, filters interface{}) ([]*model.Event, error)
	Create(ctx context.Context, event *model.Event) (*model.Event, error)
	Update(ctx context.Context, event *model.Event) (*model.Event, error)
	AddInterest(ctx context.Context, id string) (*model.Event, error)
	Delete(ctx context.Context, id string) (bool, error)
}
