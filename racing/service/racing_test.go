package service

import (
	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
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
		Name:                "race two",
		Number:              2,
		Visible:             true,
		AdvertisedStartTime: day_one,
	}

	race_four = racing.Race{
		Id:                  4,
		MeetingId:           4,
		Name:                "race four",
		Number:              4,
		Visible:             false,
		AdvertisedStartTime: day_two,
	}

	race_six = racing.Race{
		Id:                  6,
		MeetingId:           6,
		Name:                "race six",
		Number:              6,
		Visible:             true,
		AdvertisedStartTime: day_three,
	}

	race_eight = racing.Race{
		Id:                  8,
		MeetingId:           8,
		Name:                "race eight",
		Number:              8,
		Visible:             false,
		AdvertisedStartTime: day_four,
	}

	race_ten = racing.Race{
		Id:                  10,
		MeetingId:           10,
		Name:                "race ten",
		Number:              10,
		Visible:             true,
		AdvertisedStartTime: day_five,
	}
)

// Investigate "google.golang.org/grpc/test/bufconn" for unit testing the gRPC service call for
// filtering on visibility
