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
	// GetRaceById will return a single race based on the user provided id.
	GetRaceById(ctx context.Context, in *racing.GetRaceRequest) (*racing.GetRaceResponse, error)
}

// racingService implements the Racing interface.
type racingService struct {
	racesRepo db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) Racing {
	return &racingService{racesRepo}
}

// GetRaceById returns a single race retrieved using a user specified id
func (s *racingService) GetRaceById(ctx context.Context, in *racing.GetRaceRequest) (*racing.GetRaceResponse, error) {
	race, err := s.racesRepo.Get(in)
	if err != nil {
		return nil, err
	}

	race, err = AssignRaceStatusSingle(race)
	if err != nil {
		return nil, err
	}

	return &racing.GetRaceResponse{Race: race}, nil
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

// OrderRaces sorts the array of Races by the user defined order by field
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

// AssignRaceStatus compares the current time to the advertised start time of each race
// and sets the status field to either OPEN or CLOSED
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

// AssignRaceStatusSingle is a wrapper function used to call AssignRaceStatus when
// the result is a single element instead of an array of elements
func AssignRaceStatusSingle(race *racing.Race) (*racing.Race, error) {
	// make sure to test edge cases for this and justify decisions for doing it this way

	raceAsArray := []*racing.Race{race}

	result, err := AssignRaceStatus(raceAsArray)
	if err != nil {
		return nil, err
	}

	return result[0], err
}
