package db

import (
	"database/sql"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"sports/proto/sports"
)

// SportsRepo provides repository access to sports.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of sports events.
	List(filter *sports.ListSportsEventsRequestFilter) ([]*sports.SportsEvent, error)

	// Get will return a single sports event.
	Get(req *sports.GetSportsEventRequest) (*sports.SportsEvent, error)
}

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the sports repository dummy data.
func (r *sportsRepo) Init() error {
	var err error

	r.init.Do(func() {
		err = r.seed()
	})

	return err
}

// Gets and returns a single sports event based on the id provided in the request
func (r *sportsRepo) Get(req *sports.GetSportsEventRequest) (*sports.SportsEvent, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	// Builds a query that returns a sport event with the user specified id value
	query = getSportsQueries()[sport]

	// Passes the user input id into args
	args = append(args, req.Id)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// retrieve sports events returned by scanSports
	sportsEvents, err := r.scanSports(rows)
	if err != nil {
		// If error occurs, return error
		return nil, err
	} else if len(sportsEvents) == 0 {
		// If no sport event found with the id, return sql ErrNoRows error
		return nil, sql.ErrNoRows
	} else {
		// If id is valid and result is not empty, return the first sports event
		return sportsEvents[0], nil
	}
}

func (r *sportsRepo) List(filter *sports.ListSportsEventsRequestFilter) ([]*sports.SportsEvent, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportsQueries()[sportsList]

	query, args = r.applyFilter(query, filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanSports(rows)
}

func (r *sportsRepo) applyFilter(query string, filter *sports.ListSportsEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	// Filter for visible
	if filter.Visible {
		clauses = append(clauses, "visible = ?")
		args = append(args, filter.Visible)
	}

	// Filter for sport type
	if filter.Sport != nil {
		clauses = append(clauses, "sport_type IN ("+strings.Repeat("?,", len(filter.Sport)-1)+"?)")

		for _, sport_type := range filter.Sport {
			args = append(args, sport_type)
		}
	}

	// Filter for state (location)
	if filter.State != nil {
		clauses = append(clauses, "state IN ("+strings.Repeat("?,", len(filter.State)-1)+"?)")

		for _, state := range filter.State {
			args = append(args, state)
		}
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

func (m *sportsRepo) scanSports(
	rows *sql.Rows,
) ([]*sports.SportsEvent, error) {
	var sportsEvents []*sports.SportsEvent

	for rows.Next() {
		var sport sports.SportsEvent
		var advertisedStart time.Time

		if err := rows.Scan(&sport.Id, &sport.Name, &sport.SportType, &sport.State, &sport.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		sport.AdvertisedStartTime = ts

		sportsEvents = append(sportsEvents, &sport)
	}

	return sportsEvents, nil
}
