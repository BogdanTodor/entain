package db

import (
	"math/rand"
	"time"

	"syreclabs.com/go/faker"
)

// Array of strings containing samples of sports being played
var sportList = []string{
	"Hockey",
	"Baseball",
	"BasketBall",
	"Football",
}

// Creates and populates sports table with dummy data
func (r *sportsRepo) seed() error {
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, name TEXT, sport_type TEXT, state TEXT, visible INTEGER, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO sports(id, name, sport_type, state, visible, advertised_start_time) VALUES (?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Team().Name(),
				sportList[rand.Intn(3)],
				faker.Team().State(),
				faker.Number().Between(0, 1),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}
