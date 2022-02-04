package service

import (
	"testing"
	"time"

	"git.neds.sh/matty/entain/racing/proto/racing"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Creating test entities
var (
	// Creating timestamps for Race advertised_state_time field\
	day_one   = timestamppb.New(time.Now().AddDate(0, 0, -4))
	day_two   = timestamppb.New(time.Now().AddDate(0, 0, -3))
	day_three = timestamppb.New(time.Now().AddDate(0, 0, -2))
	day_four  = timestamppb.New(time.Now().AddDate(0, 0, -1))
	day_five  = timestamppb.New(time.Now())

	// Creating five separate races to perform tests with
	race_two = racing.Race{
		Id:                  2,
		MeetingId:           2,
		Name:                "two",
		Number:              2,
		Visible:             true,
		AdvertisedStartTime: day_one,
	}

	race_four = racing.Race{
		Id:                  4,
		MeetingId:           4,
		Name:                "four",
		Number:              4,
		Visible:             false,
		AdvertisedStartTime: day_two,
	}

	race_six = racing.Race{
		Id:                  6,
		MeetingId:           6,
		Name:                "six",
		Number:              6,
		Visible:             true,
		AdvertisedStartTime: day_three,
	}

	race_eight = racing.Race{
		Id:                  8,
		MeetingId:           8,
		Name:                "eight",
		Number:              8,
		Visible:             false,
		AdvertisedStartTime: day_four,
	}

	race_ten = racing.Race{
		Id:                  10,
		MeetingId:           10,
		Name:                "ten",
		Number:              10,
		Visible:             true,
		AdvertisedStartTime: day_five,
	}
)

// Investigate "google.golang.org/grpc/test/bufconn" for unit testing the gRPC service call for
// filtering on visibility

// Test order by functionality with different data types (int64, string and timestamp)
func TestOrderRaces(t *testing.T) {
	type inputArgs struct {
		races   []*racing.Race
		orderBy string
	}
	tests := []struct {
		description string
		input       inputArgs
		expect      []*racing.Race
		err         error
	}{
		{
			description: "sorts array by id",
			input: inputArgs{
				races: []*racing.Race{
					&race_two,
					&race_eight,
					&race_six,
					&race_four,
					&race_ten,
				},
				orderBy: "id",
			},
			expect: []*racing.Race{
				&race_two,
				&race_four,
				&race_six,
				&race_eight,
				&race_ten,
			},
			err: nil,
		},
		{
			description: "sorts array by name",
			input: inputArgs{
				races: []*racing.Race{
					&race_two,
					&race_four,
					&race_six,
					&race_eight,
					&race_ten,
				},
				orderBy: "name",
			},
			expect: []*racing.Race{
				&race_eight,
				&race_four,
				&race_six,
				&race_ten,
				&race_two,
			},
			err: nil,
		},
		{
			description: "sorts array by advertised_start_time",
			input: inputArgs{
				races: []*racing.Race{
					&race_eight,
					&race_four,
					&race_six,
					&race_ten,
					&race_two,
				},
				orderBy: "advertised_start_time",
			},
			expect: []*racing.Race{
				&race_two,
				&race_four,
				&race_six,
				&race_eight,
				&race_ten,
			},
			err: nil,
		},
		{
			description: "handle empty input array",
			input: inputArgs{
				races:   []*racing.Race{},
				orderBy: "id",
			},
			expect: []*racing.Race{},
			err:    nil,
		},
		{
			description: "handle input array with only one element",
			input: inputArgs{
				races: []*racing.Race{
					&race_two,
				},
				orderBy: "id",
			},
			expect: []*racing.Race{
				&race_two,
			},
			err: nil,
		},
		{
			description: "handles invalid user input order by term",
			input: inputArgs{
				races: []*racing.Race{
					&race_eight,
					&race_four,
					&race_six,
					&race_ten,
					&race_two,
				},
				orderBy: "meating_id",
			},
			expect: []*racing.Race{
				&race_two,
				&race_four,
				&race_six,
				&race_eight,
				&race_ten,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := OrderRaces(tt.input.races, tt.input.orderBy)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.err, err)
		})
	}
}
