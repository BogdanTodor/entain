package db

const (
	sportsList = "list"
	sport      = "get"
)

func getSportsQueries() map[string]string {
	return map[string]string{
		// Query database to return all sports events
		sportsList: `
			SELECT 
				id, 
				name, 
				sport_type, 
				state, 
				visible, 
				advertised_start_time 
			FROM sports
		`,
		// Query database for a single sports event based on the provided id
		sport: `
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
		`,
	}
}
