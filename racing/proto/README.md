## Racing Proto

### API

#### Race
* **ID**: The unique identifier for the race (int64)
* **MeetingID**: The unique identifier for the races meeting (int64)
* **Name**: The name given to the race (string)
* **Number**: The number of the race (int64)
* **Visible**: Whether the race is visible (bool)
* **AdvertisedStartTime**: The time the race is advertised to run (Timestamp)

#### RPCs
```ListRaces(ListRacesRequest) ListRacesResponse```

#### Requests
* **ListRacesRequest**: Request which takes ```ListRacesRequestFilter``` and ```order_by``` as inputs 

#### Responses
* **ListRacesResponse**: The response containing all races that match the filtering criteria

#### Filters
* **ListRacesRequestFilter**: Input into ListRaceRequest that filters the races