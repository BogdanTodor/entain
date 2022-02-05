package db

const (
	racesList = "list"
	race      = "get"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		// Query database to return all races
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
		`,
		// Query database for a single race based on the provided id
		race: `
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
		`,
	}
}
