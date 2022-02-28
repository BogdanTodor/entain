package db

import (
	"regexp"
	"testing"
	"time"

	"git.neds.sh/matty/entain/racing/proto/racing"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Create test entities
var (

	// Create timestamps for Race advertised_start_time field
	day_one   = time.Now().AddDate(0, 0, -2)
	day_two   = time.Now().AddDate(0, 0, -1)
	day_three = time.Now().AddDate(0, 0, 1)

	// Create three separate races to perform tests with
	race_two = racing.Race{
		Id:                  2,
		MeetingId:           2,
		Name:                "two",
		Number:              2,
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(day_one),
		Status:              "",
	}

	race_four = racing.Race{
		Id:                  4,
		MeetingId:           4,
		Name:                "four",
		Number:              4,
		Visible:             false,
		AdvertisedStartTime: timestamppb.New(day_two),
		Status:              "",
	}

	race_six = racing.Race{
		Id:                  6,
		MeetingId:           6,
		Name:                "six",
		Number:              6,
		Visible:             true,
		AdvertisedStartTime: timestamppb.New(day_three),
		Status:              "",
	}
)

func TestGet(t *testing.T) {
	db, mock, err1 := sqlmock.New()
	repo := NewRacesRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					meeting_id,
					name,
					number,
					visible,
					advertised_start_time
				FROM races
				WHERE id = ?
				LIMIT 1
			`
	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "meeting_id", "name", "number", "visible", "advertised_start_time"}).
		AddRow(race_two.Id, race_two.MeetingId, race_two.Name, race_two.Number, race_two.Visible, day_one)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(race_two.Id).WillReturnRows(rows)

	// Query using valid id
	request := racing.GetRaceRequest{Id: race_two.Id}
	response, err := repo.Get(&request)

	// Verify results are as expected
	assert.NotNil(t, response)
	assert.Equal(t, race_two.Id, response.Id)
	assert.NoError(t, err)
}

func TestGetError(t *testing.T) {
	db, mock, err1 := sqlmock.New()
	repo := NewRacesRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					meeting_id,
					name,
					number,
					visible,
					advertised_start_time
				FROM races
				WHERE id = ?
				LIMIT 1
			`
	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "meeting_id", "name", "number", "visible", "advertised_start_time"}).
		AddRow(race_two.Id, race_two.MeetingId, race_two.Name, race_two.Number, race_two.Visible, day_one)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(race_two.Id).WillReturnRows(rows)

	// Query using invalid id
	request := racing.GetRaceRequest{Id: race_four.Id}
	response, err := repo.Get(&request)

	// Verify results are as expected
	assert.Empty(t, response)
	assert.Error(t, err)
}

func TestListAndFilterVisible(t *testing.T) {
	db, mock, err1 := sqlmock.New()
	repo := NewRacesRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					meeting_id,
					name,
					number,
					visible,
					advertised_start_time
				FROM races
				WHERE visible = ?
			`
	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "meeting_id", "name", "number", "visible", "advertised_start_time"}).
		AddRow(race_two.Id, race_two.MeetingId, race_two.Name, race_two.Number, race_two.Visible, day_one).
		AddRow(race_six.Id, race_six.MeetingId, race_six.Name, race_six.Number, race_six.Visible, day_three)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(true).WillReturnRows(rows)

	// Query using visible filter
	filter := racing.ListRacesRequestFilter{Visible: true}
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
	repo := NewRacesRepo(db)
	if err1 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err1)
	}
	defer db.Close()

	// Create query
	query := `
				SELECT
					id,
					meeting_id,
					name,
					number,
					visible,
					advertised_start_time
				FROM races
			`
	// Prepare mock result data and mock the query
	rows := sqlmock.NewRows([]string{"id", "meeting_id", "name", "number", "visible", "advertised_start_time"}).
		AddRow(race_two.Id, race_two.MeetingId, race_two.Name, race_two.Number, race_two.Visible, day_one).
		AddRow(race_four.Id, race_four.MeetingId, race_four.Name, race_four.Number, race_four.Visible, day_two).
		AddRow(race_six.Id, race_six.MeetingId, race_six.Name, race_six.Number, race_six.Visible, day_three)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs().WillReturnRows(rows)

	// Query using empty filter
	filter := racing.ListRacesRequestFilter{}
	response, err := repo.List(&filter)

	// Assert result is as expected
	assert.NotNil(t, response)
	assert.Equal(t, len(response), 3)
	assert.NoError(t, err)
}
