package service

import (
	"sort"
	"time"

	"sports/db"
	"sports/proto/sports"

	"golang.org/x/net/context"
)

type SportsEvents interface {
	// ListSportsEvents will return a collection of sports events.
	ListSportsEvents(ctx context.Context, in *sports.ListSportsEventsRequest) (*sports.ListSportsEventsResponse, error)
	// GetSportsEventById will return a single sports event based on the provided id
	GetSportsEventById(ctx context.Context, in *sports.GetSportsEventRequest) (*sports.GetSportsEventResponse, error)
}

// sportsService implements the SportsEvents interface.
type sportsService struct {
	sportsRepo db.SportsRepo
}

// NewSportsService instantiates and returns a new sportsService.
func NewSportsService(sportsRepo db.SportsRepo) SportsEvents {
	return &sportsService{sportsRepo}
}

// Return a single sports event retrieved using a user specified id
func (s *sportsService) GetSportsEventById(ctx context.Context, in *sports.GetSportsEventRequest) (*sports.GetSportsEventResponse, error) {
	// Retrieve a single row from the database containing the sports event by id
	sportEvent, err := s.sportsRepo.Get(in)
	if err != nil {
		return nil, err
	}

	// Assign a status of either OPEN or CLOSED to the sports event
	sportEvent, err = AssignSportsEventStatusSingle(sportEvent)
	if err != nil {
		return nil, err
	}

	return &sports.GetSportsEventResponse{SportsEvent: sportEvent}, nil
}

// Returns an array of sports events
func (s *sportsService) ListSportsEvents(ctx context.Context, in *sports.ListSportsEventsRequest) (*sports.ListSportsEventsResponse, error) {
	// Retrieves all rows that meet the criteria of the filter from the database table
	sportsEvents, err := s.sportsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	// Assign a status of either OPEN or CLOSED to each sports event
	sportsEvents, err = AssignSportsEventsStatus(sportsEvents)
	if err != nil {
		return nil, err
	}

	// Order the sports events by the user specified field (defaults to advertised time)
	sportsEvents, err = OrderSportsEvents(sportsEvents, in.OrderBy)
	if err != nil {
		return nil, err
	}

	return &sports.ListSportsEventsResponse{SportsEvents: sportsEvents}, nil
}

// Returns a sorted array of sports events based on the user defined order by field
func OrderSportsEvents(sportsEvents []*sports.SportsEvent, orderedBy string) ([]*sports.SportsEvent, error) {
	// Sorts array by any of the below valid fields
	// * defaults to advertised start time if not specified or invalid field is provided
	sort.Slice(sportsEvents, func(i, j int) bool {
		switch orderedBy {
		case "advertised_start_time":
			return sportsEvents[i].GetAdvertisedStartTime().GetSeconds() < sportsEvents[j].GetAdvertisedStartTime().GetSeconds()
		case "id":
			return sportsEvents[i].GetId() < sportsEvents[j].GetId()
		case "name":
			return sportsEvents[i].GetName() < sportsEvents[j].GetName()
		case "sport_type":
			return sportsEvents[i].GetSportType() < sportsEvents[j].GetSportType()
		case "state":
			return sportsEvents[i].GetState() < sportsEvents[j].GetState()
		default:
			return sportsEvents[i].GetAdvertisedStartTime().GetSeconds() < sportsEvents[j].GetAdvertisedStartTime().GetSeconds()
		}
	})
	return sportsEvents, nil
}

// Assigns the status field to either OPEN or CLOSED and returns the resulting array
func AssignSportsEventsStatus(sportsEvents []*sports.SportsEvent) ([]*sports.SportsEvent, error) {
	currentTime := time.Now()

	for _, sportEvent := range sportsEvents {
		// Convert advertised start time to standard Go time
		advertisedStart := sportEvent.GetAdvertisedStartTime().AsTime()

		// Compare sports event start time to current time and assign status accordingly
		if advertisedStart.Before(currentTime) {
			sportEvent.Status = "CLOSED"
		} else {
			sportEvent.Status = "OPEN"
		}
	}
	return sportsEvents, nil
}

// Wrapper function to call AssignSportsEventsStatus when the input and result are single element not arrays
func AssignSportsEventStatusSingle(sportsEvent *sports.SportsEvent) (*sports.SportsEvent, error) {
	// Store single sports event in array to pass to AssignSportsEventsStatus
	sportsEventAsArray := []*sports.SportsEvent{sportsEvent}

	result, err := AssignSportsEventsStatus(sportsEventAsArray)
	if err != nil {
		return nil, err
	}

	return result[0], err
}
