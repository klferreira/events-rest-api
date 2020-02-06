package event

import (
	"context"

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
	return s.repo.Create(ctx, event)
}

func (s *service) Update(ctx context.Context, event *model.Event) (*model.Event, error) {
	return s.repo.Update(ctx, event)
}

func (s *service) AddInterest(ctx context.Context, id string) (*model.Event, error) {
	event, err := s.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	event.Interested++

	event, err = s.repo.Update(ctx, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *service) Delete(ctx context.Context, id string) (bool, error) {
	return s.repo.Delete(ctx, id)
}
