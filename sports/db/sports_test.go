package db

import (
	"regexp"
	"testing"
	"time"

	"sports/proto/sports"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Creating test entities
var (

	// Creating timestamps for SportsEvent advertised_state_time field
	day_one   = time.Now().AddDate(0, 0, -2)
	day_two   = time.Now().AddDate(0, 0, -1)
	day_three = time.Now().AddDate(0, 0, 1)
	day_four  = time.Now().AddDate(0, 0, 2)

	// Creating four separate sports events to perform tests with
	event_one = sports.SportsEvent{
		Id:                  1,
		Name:                "Hockey Event 1",
		SportType:           "Hockey",
		State:               "New York",
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(day_one),
		Status:              "",
	}

	event_two = sports.SportsEvent{
		Id:                  2,
		Name:                "MLB finals",
		SportType:           "Baseball",
		State:               "New York",
		Visible:             false,
		AdvertisedStartTime: timestamppb.New(day_two),
		Status:              "",
	}

	event_three = sports.SportsEvent{
		Id:                  3,
		Name:                "NBA finals",
		SportType:           "Basketball",
		State:               "California",
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(day_three),
		Status:              "",
	}

	event_four = sports.SportsEvent{
		Id:                  4,
		Name:                "Superbowl",
		SportType:           "Football",
		State:               "Texas",
		Visible:             false,
		AdvertisedStartTime: timestamppb.New(day_four),
		Status:              "",
	}
)

func TestGet(t *testing.T) {
	db, mock, err1 := sqlmock.New()
	repo := NewSportsRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					name,
					sport_type,
					state,
					visible,
					advertised_start_time
				FROM sports
				WHERE id = ?
				LIMIT 1
			`
	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "name", "sport_type", "state", "visible", "advertised_start_time"}).
		AddRow(event_one.Id, event_one.Name, event_one.SportType, event_one.State, event_one.Visible, day_one)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(event_one.Id).WillReturnRows(rows)

	// Query using valid id
	request := sports.GetSportsEventRequest{Id: event_one.Id}
	response, err := repo.Get(&request)

	// Verify results are as expected
	assert.NotNil(t, response)
	assert.Equal(t, event_one.Id, response.Id)
	assert.NoError(t, err)
}

func TestGetError(t *testing.T) {
	db, mock, err1 := sqlmock.New()
	repo := NewSportsRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					name,
					sport_type,
					state,
					visible,
					advertised_start_time
				FROM sports
				WHERE id = ?
				LIMIT 1
			`
	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "name", "sport_type", "state", "visible", "advertised_start_time"}).
		AddRow(event_one.Id, event_one.Name, event_one.SportType, event_one.State, event_one.Visible, day_one)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(event_one.Id).WillReturnRows(rows)

	// Query using invalid id
	request := sports.GetSportsEventRequest{Id: event_two.Id}
	response, err := repo.Get(&request)
	assert.Empty(t, response)
	assert.Error(t, err)
}

func TestListAndFilterVisible(t *testing.T) {
	db, mock, err1 := sqlmock.New()
	repo := NewSportsRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					name,
					sport_type,
					state,
					visible,
					advertised_start_time
				FROM sports
				WHERE visible = ?
			`

	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "name", "sport_type", "state", "visible", "advertised_start_time"}).
		AddRow(event_one.Id, event_one.Name, event_one.SportType, event_one.State, event_one.Visible, day_one).
		AddRow(event_three.Id, event_three.Name, event_three.SportType, event_three.State, event_three.Visible, day_three)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(true).WillReturnRows(rows)

	// Query using visible filter
	filter := sports.ListSportsEventsRequestFilter{Visible: true}
	response, err := repo.List(&filter)

	// Assert result is as expected
	assert.NotNil(t, response)
	assert.Equal(t, 2, len(response))
	assert.Equal(t, response[0].Visible, true)
	assert.Equal(t, response[1].Visible, true)
	assert.NoError(t, err)
}

func TestListEmptyFilter(t *testing.T) {
	db, mock, err1 := sqlmock.New()
	repo := NewSportsRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					name,
					sport_type,
					state,
					visible,
					advertised_start_time
				FROM sports
			`
	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "name", "sport_type", "state", "visible", "advertised_start_time"}).
		AddRow(event_one.Id, event_one.Name, event_one.SportType, event_one.State, event_one.Visible, day_one).
		AddRow(event_two.Id, event_two.Name, event_two.SportType, event_two.State, event_two.Visible, day_two).
		AddRow(event_three.Id, event_three.Name, event_three.SportType, event_three.State, event_three.Visible, day_one).
		AddRow(event_four.Id, event_four.Name, event_four.SportType, event_four.State, event_four.Visible, day_two)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs().WillReturnRows(rows)

	// Query using empty filter
	filter := sports.ListSportsEventsRequestFilter{}
	response, err := repo.List(&filter)

	// Assert result is as expected
	assert.NotNil(t, response)
	assert.Equal(t, len(response), 4)
	assert.NoError(t, err)
}
