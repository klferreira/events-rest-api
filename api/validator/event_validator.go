package validator

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/klferreira/events-rest-api/internal/model"
)

func getEventFromRequestBody(r io.ReadCloser) (*model.Event, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	event := &model.Event{}

	err = json.Unmarshal(body, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func ValidateEventCreateRequest(req *http.Request) (*model.Event, error) {
	event, err := getEventFromRequestBody(req.Body)
	if err != nil {
		return nil, errors.New("could not read/parse event input")
	}

	if len(event.Name) < 3 {
		return nil, errors.New("event name has to be at least 3 chars long")
	}

	if len(event.Place) < 5 {
		return nil, errors.New("event place should be at least 5 chars long")
	}

	if len(event.Sessions) == 0 {
		return nil, errors.New("event must have at least 1 session")
	}

	return event, nil
}

func ValidateEventUpdateRequest(req *http.Request) (*model.Event, error) {
	event, err := ValidateEventCreateRequest(req)
	if err != nil {
		return nil, err
	}

	if !event.ID.Valid() {
		return nil, errors.New("cannot update event with nil id")
	}

	return event, nil
}
