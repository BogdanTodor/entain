package service

import (
	"sort"
	"time"

	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"golang.org/x/net/context"
)

type Racing interface {
	// ListRaces will return a collection of races.
	ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error)
}

// racingService implements the Racing interface.
type racingService struct {
	racesRepo db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) Racing {
	return &racingService{racesRepo}
}

func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {
	races, err := s.racesRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	races, err = AssignRaceStatus(races)
	if err != nil {
		return nil, err
	}

	races, err = OrderRaces(races, in.OrderBy)
	if err != nil {
		return nil, err
	}

	return &racing.ListRacesResponse{Races: races}, nil
}

func OrderRaces(races []*racing.Race, orderedBy string) ([]*racing.Race, error) {

	// Sorts the races array using sort.Slice based on the orderedBy user input
	// where in the case of no input or invalid inputs, advertised_start_time is
	// the default case

	sort.Slice(races, func(i, j int) bool {
		switch orderedBy {
		case "advertised_start_time":
			return races[i].GetAdvertisedStartTime().GetSeconds() < races[j].GetAdvertisedStartTime().GetSeconds()
		case "id":
			return races[i].GetId() < races[j].GetId()
		case "meeting_id":
			return races[i].GetMeetingId() < races[j].GetMeetingId()
		case "name":
			return races[i].GetName() < races[j].GetName()
		case "number":
			return races[i].GetNumber() < races[j].GetNumber()
		default:
			return races[i].GetAdvertisedStartTime().GetSeconds() < races[j].GetAdvertisedStartTime().GetSeconds()
		}
	})
	return races, nil
}

func AssignRaceStatus(races []*racing.Race) ([]*racing.Race, error) {
	// Get the current timestamp
	currentTime := time.Now()

	// Iterate over the races and compare the current time to the advertised start time
	// to assign the status
	for _, race := range races {
		// Convert advertised start time to standard Go time
		advertisedStart := race.GetAdvertisedStartTime().AsTime()

		if advertisedStart.Before(currentTime) {
			race.Status = "CLOSED"
		} else {
			race.Status = "OPEN"
		}
	}
	return races, nil
}
