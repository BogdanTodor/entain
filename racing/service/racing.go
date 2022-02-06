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

// Returns a single race retrieved using a user specified id
func (s *racingService) GetRaceById(ctx context.Context, in *racing.GetRaceRequest) (*racing.GetRaceResponse, error) {
	// Retrieve a single row from the database containing the race by id
	race, err := s.racesRepo.Get(in)
	if err != nil {
		return nil, err
	}

	// Assign a status of either OPEN or CLOSED to the race
	race, err = AssignRaceStatusSingle(race)
	if err != nil {
		return nil, err
	}

	return &racing.GetRaceResponse{Race: race}, nil
}

// Returns an array of races
func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {
	// Retrieves all rows that meet the criteria of the filter from the database table
	races, err := s.racesRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	// Assign a status of either OPEN or CLOSED to each race
	races, err = AssignRaceStatus(races)
	if err != nil {
		return nil, err
	}

	// Order the races by the user specified field (defaults to advertised time)
	races, err = OrderRaces(races, in.OrderBy)
	if err != nil {
		return nil, err
	}

	return &racing.ListRacesResponse{Races: races}, nil
}

// Returns a sorted array of races based on the user defined order by field
func OrderRaces(races []*racing.Race, orderedBy string) ([]*racing.Race, error) {
	// Sorts array by any of the below valid fields
	// * defaults to advertised start time if not specified or invalid field is provided
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

// Assigns the status field to either OPEN or CLOSED and returns the resulting array
func AssignRaceStatus(races []*racing.Race) ([]*racing.Race, error) {
	currentTime := time.Now()

	for _, race := range races {
		// Convert advertised start time to standard Go time
		advertisedStart := race.GetAdvertisedStartTime().AsTime()

		// Compare race start time to current time and assign status accordingly
		race.Status = "OPEN"
		if advertisedStart.Before(currentTime) {
			race.Status = "CLOSED"
		}
	}
	return races, nil
}

// Wrapper function to call AssignRaceStatus when the input and result are single element not arrays
func AssignRaceStatusSingle(race *racing.Race) (*racing.Race, error) {
	// Store single race in array to pass to AssignRaceStatus
	raceAsArray := []*racing.Race{race}

	result, err := AssignRaceStatus(raceAsArray)
	if err != nil {
		return nil, err
	}

	return result[0], err
}
