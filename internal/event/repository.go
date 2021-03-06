package event

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"github.com/klferreira/events-rest-api/internal/model"
	"github.com/klferreira/events-rest-api/pkg/mongo"
)

type repository struct {
	db mongo.Client
}

func NewRepository(db mongo.Client) Repository {
	return &repository{db}
}

func (r *repository) Get(ctx context.Context, filters interface{}) (events []*model.Event, err error) {
	err = r.db.FindAll("events", filters, &events)
	return
}

func (r *repository) Find(ctx context.Context, id string) (*model.Event, error) {
	event := &model.Event{}

	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := r.db.FindOne("events", selector, event); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repository) Create(ctx context.Context, event *model.Event) (*model.Event, error) {
	event.ID = bson.NewObjectId()

	err := r.db.Insert("events", event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repository) Update(ctx context.Context, event *model.Event) (*model.Event, error) {
	selector := bson.M{"_id": event.ID}
	err := r.db.Update("events", selector, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repository) Delete(ctx context.Context, id string) (bool, error) {
	err := r.db.DeleteOne("events", id)
	if err != nil {
		return false, err
	}

	return true, nil
}
