package service

import (
	"testing"
	"time"

	"sports/proto/sports"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Creating test entities
var (

	// Creating timestamps for sports events advertised_state_time field
	day_one   = timestamppb.New(time.Now().AddDate(0, 0, -2))
	day_two   = timestamppb.New(time.Now().AddDate(0, 0, -1))
	day_three = timestamppb.New(time.Now().AddDate(0, 0, 1))
	day_four  = timestamppb.New(time.Now().AddDate(0, 0, 2))

	// Creating four separate sports events to perform tests with
	event_one = sports.SportsEvent{
		Id:                  1,
		Name:                "Hockey Event 1",
		SportType:           "Hockey",
		State:               "New York",
		Visible:             true,
		AdvertisedStartTime: day_one,
		Status:              "",
	}

	event_two = sports.SportsEvent{
		Id:                  2,
		Name:                "MLB finals",
		SportType:           "Baseball",
		State:               "New York",
		Visible:             false,
		AdvertisedStartTime: day_two,
		Status:              "",
	}

	event_three = sports.SportsEvent{
		Id:                  3,
		Name:                "NBA finals",
		SportType:           "Basketball",
		State:               "California",
		Visible:             true,
		AdvertisedStartTime: day_three,
		Status:              "",
	}

	event_four = sports.SportsEvent{
		Id:                  4,
		Name:                "Superbowl",
		SportType:           "Football",
		State:               "Texas",
		Visible:             false,
		AdvertisedStartTime: day_four,
		Status:              "",
	}
)

func TestOrderSportsEvents(t *testing.T) {
	type inputArgs struct {
		sportsEvents []*sports.SportsEvent
		orderBy      string
	}
	tests := []struct {
		description string
		input       inputArgs
		expect      []*sports.SportsEvent
		err         error
	}{
		{
			description: "sorts array by id",
			input: inputArgs{
				sportsEvents: []*sports.SportsEvent{
					&event_one,
					&event_three,
					&event_four,
					&event_two,
				},
				orderBy: "id",
			},
			expect: []*sports.SportsEvent{
				&event_one,
				&event_two,
				&event_three,
				&event_four,
			},
			err: nil,
		},
		{
			description: "sorts array by name",
			input: inputArgs{
				sportsEvents: []*sports.SportsEvent{
					&event_one,
					&event_two,
					&event_three,
					&event_four,
				},
				orderBy: "name",
			},
			expect: []*sports.SportsEvent{
				&event_one,
				&event_two,
				&event_three,
				&event_four,
			},
			err: nil,
		},
		{
			description: "sorts array by advertised_start_time",
			input: inputArgs{
				sportsEvents: []*sports.SportsEvent{
					&event_two,
					&event_three,
					&event_four,
					&event_one,
				},
				orderBy: "advertised_start_time",
			},
			expect: []*sports.SportsEvent{
				&event_one,
				&event_two,
				&event_three,
				&event_four,
			},
			err: nil,
		},
		{
			description: "handle empty input array",
			input: inputArgs{
				sportsEvents: []*sports.SportsEvent{},
				orderBy:      "id",
			},
			expect: []*sports.SportsEvent{},
			err:    nil,
		},
		{
			description: "handle input array with only one element",
			input: inputArgs{
				sportsEvents: []*sports.SportsEvent{
					&event_one,
				},
				orderBy: "id",
			},
			expect: []*sports.SportsEvent{
				&event_one,
			},
			err: nil,
		},
		{
			description: "handles invalid user input order by term",
			input: inputArgs{
				sportsEvents: []*sports.SportsEvent{
					&event_four,
					&event_two,
					&event_three,
					&event_one,
				},
				orderBy: "nme",
			},
			expect: []*sports.SportsEvent{
				&event_one,
				&event_two,
				&event_three,
				&event_four,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := OrderSportsEvents(tt.input.sportsEvents, tt.input.orderBy)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAssignSportsEventsStatus(t *testing.T) {
	tests := []struct {
		description string
		input       []*sports.SportsEvent
		expect      []string
		err         error
	}{
		{
			description: "Assigns status field correctly",
			input: []*sports.SportsEvent{
				&event_one,
				&event_two,
				&event_three,
				&event_four,
			},
			expect: []string{
				"CLOSED",
				"CLOSED",
				"OPEN",
				"OPEN",
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := AssignSportsEventsStatus(tt.input)
			// Create slice to store the statuses of the sports events returned by the
			// AssignSportsEventsStatus function and compare to expected statuses
			actualStatuses := []string{}

			for sportsEventIndex := range actual {
				actualStatuses = append(actualStatuses, actual[sportsEventIndex].Status)
			}

			assert.Equal(t, tt.expect, actualStatuses)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestAssignSportsEventStatusSingle(t *testing.T) {
	tests := []struct {
		description string
		input       *sports.SportsEvent
		expect      string
		err         error
	}{
		{
			description: "Assigns status CLOSED to sports event from the past",
			input:       &event_one,
			expect:      "CLOSED",
			err:         nil,
		},
		{
			description: "Assigns status OPEN to sports event in the future",
			input:       &event_four,
			expect:      "OPEN",
			err:         nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := AssignSportsEventStatusSingle(tt.input)
			assert.Equal(t, tt.expect, actual.Status)
			assert.Equal(t, tt.err, err)
		})
	}
}
