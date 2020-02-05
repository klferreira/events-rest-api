package event

import (
	"context"
	"errors"

	"github.com/klferreira/events-rest-api/internal/model"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Fetch(ctx context.Context, filters interface{}) ([]*model.Event, error) {
	events, err := s.repo.Get(ctx, filters)
	if err == nil && events == nil {
		return []*model.Event{}, nil
	}

	return events, err
}

func (s *service) Create(ctx context.Context, event *model.Event) (*model.Event, error) {
	if len(event.Name) < 5 {
		return nil, errors.New("event name should be at least 5 chars long")
	}

	return s.repo.Create(ctx, event)
}

func (s *service) Update(ctx context.Context, event *model.Event) (*model.Event, error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, event *model.Event) (*model.Event, error) {
	return nil, nil
}
